/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
	"github.com/kurtosis-tech/kurtosis-client/golang/kurtosis_core_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/services"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_docker_api"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_consts"
	"github.com/kurtosis-tech/minimal-grpc-server/server"
	"github.com/palantir/stacktrace"
	"google.golang.org/grpc"
	"time"
)

const (
	grpcServerStopGracePeriod = 5 * time.Second
)

type LambdaExecutor struct {
	apiContainerSocket        string
	serializedCustomParamsStr string
	configurator              LambdaConfigurator
}

func NewLambdaExecutor(apiContainerSocket string, serializedCustomParamsStr string, configurator LambdaConfigurator) *LambdaExecutor {
	return &LambdaExecutor{apiContainerSocket: apiContainerSocket, serializedCustomParamsStr: serializedCustomParamsStr, configurator: configurator}
}

func (executor LambdaExecutor) Run() error {

	lambda, err := executor.configurator.ParseParamsAndCreateLambda(executor.serializedCustomParamsStr)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing the serialized custom params and creating the Lambda")
	}

	if executor.apiContainerSocket == "" {
		return stacktrace.NewError("The executor's field 'apiContainerSocket' was unexpectedly empty")
	}
	// TODO SECURITY: Use HTTPS to verify we're hitting the correct API container
	conn, err := grpc.Dial(executor.apiContainerSocket, grpc.WithInsecure())
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the API container at '%v'", executor.apiContainerSocket)
	}

	apiClient := kurtosis_core_rpc_api_bindings.NewApiContainerServiceClient(conn)
	networkCtx := networks.NewNetworkContext(
		apiClient,
		map[services.FilesArtifactID]string{},
		kurtosis_lambda_docker_api.ExecutionVolumeMountpoint,
	)

	lambdaServiceServer := NewLambdaServiceServer(lambda, networkCtx)
	lambdaServiceRegistrationFunc := func(grpcServer *grpc.Server) {
		kurtosis_lambda_rpc_api_bindings.RegisterLambdaServiceServer(grpcServer, lambdaServiceServer)
	}

	lambdaServer := server.NewMinimalGRPCServer(
		kurtosis_lambda_rpc_api_consts.ListenPort,
		kurtosis_lambda_rpc_api_consts.ListenProtocol,
		grpcServerStopGracePeriod,
		[]func(desc *grpc.Server){
			lambdaServiceRegistrationFunc,
		},
	)
	if err := lambdaServer.Run(); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the Lambda server")
	}

	return nil
}

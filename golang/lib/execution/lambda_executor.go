/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
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
	logLevelStr               string
	serializedCustomParamsStr string
	configurator              LambdaConfigurator
}

func NewLambdaExecutor(apiContainerSocket string, logLevelStr string, serializedCustomParamsStr string, configurator LambdaConfigurator) *LambdaExecutor {
	return &LambdaExecutor{apiContainerSocket: apiContainerSocket, logLevelStr: logLevelStr, serializedCustomParamsStr: serializedCustomParamsStr, configurator: configurator}
}

func (executor LambdaExecutor) Run() error {
	if err := executor.configurator.SetLogLevel(executor.logLevelStr); err != nil {
		return stacktrace.Propagate(err, "An error occurred setting the loglevel before running the lambda")
	}

	lambda, err := executor.configurator.ParseParamsAndCreateLambda(executor.serializedCustomParamsStr)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing the serialized custom params and creating the lambda")
	}

	var lambdaServiceClient kurtosis_lambda_rpc_api_bindings.LambdaServiceClient = nil
	if executor.apiContainerSocket != "" {
		// TODO SECURITY: Use HTTPS to ensure we're connecting to the real Lamba API servers
		conn, err := grpc.Dial(executor.apiContainerSocket, grpc.WithInsecure())
		if err != nil {
			return stacktrace.Propagate(
				err,
				"An error occurred creating a connection to the Kurtosis API server at '%v'",
				executor.apiContainerSocket,
			)
		}
		defer conn.Close()

		lambdaServiceClient = kurtosis_lambda_rpc_api_bindings.NewLambdaServiceClient(conn)
	}

	lambdaServiceServer := NewLambdaServiceServer(lambda, lambdaServiceClient)
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
		return stacktrace.Propagate(err, "An error occurred running the lambda server")
	}

	return nil
}

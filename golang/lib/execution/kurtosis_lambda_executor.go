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
	"os"
	"time"
)

const (
	grpcServerStopGracePeriod = 5 * time.Second
)

type KurtosisLambdaExecutor struct {
	configurator KurtosisLambdaConfigurator
}

func NewKurtosisLambdaExecutor(configurator KurtosisLambdaConfigurator) *KurtosisLambdaExecutor {
	return &KurtosisLambdaExecutor{configurator: configurator}
}

func (executor KurtosisLambdaExecutor) Run() error {

	serializedCustomParams, err := getEnvVar(kurtosis_lambda_docker_api.SerializedCustomParamsEnvVar, "the serialized custom params that the Lambda will consume")
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred when trying to get the serialized custom params environment variable")
	}

	lambda, err := executor.configurator.ParseParamsAndCreateKurtosisLambda(serializedCustomParams)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing the serialized custom params and creating the Lambda")
	}

	apiContainerSocket, err := getEnvVar(kurtosis_lambda_docker_api.ApiContainerSocketEnvVar, "the socket value used in API container connection")
    if err != nil {
		return stacktrace.Propagate(err, "An error occurred when trying to get the API container socket environment variable")
	}

	// TODO SECURITY: Use HTTPS to verify we're hitting the correct API container
	conn, err := grpc.Dial(apiContainerSocket, grpc.WithInsecure())
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the API container at '%v'", apiContainerSocket)
	}

	apiClient := kurtosis_core_rpc_api_bindings.NewApiContainerServiceClient(conn)
	networkCtx := networks.NewNetworkContext(
		apiClient,
		map[services.FilesArtifactID]string{},
		kurtosis_lambda_docker_api.ExecutionVolumeMountpoint,
	)

	lambdaServiceServer := NewKurtosisLambdaServiceServer(lambda, networkCtx)
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

// ====================================================================================================
//                                       Private helper functions
// ====================================================================================================
func getEnvVar(envVarName string, envVarDescription string) (string, error) {
	envVarValue, found := os.LookupEnv(envVarName)

	if !found {
		return "", stacktrace.NewError("Expected an '%v' environment variable containing '%v', but none was found", envVarName, envVarDescription)
	}
	if envVarValue == "" {
		return "", stacktrace.NewError("The '%v' serialized custom params environment variable was defined, but is emptystring", envVarName)
	}

	return envVarValue, nil
}

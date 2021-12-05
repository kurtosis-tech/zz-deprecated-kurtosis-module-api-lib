/*
 *    Copyright 2021 Kurtosis Technologies Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package execution

import (
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/kurtosis_core_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/kurtosis_module_docker_api"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/kurtosis_module_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/kurtosis_module_rpc_api_consts"
	grpc_server "github.com/kurtosis-tech/minimal-grpc-server/golang/server"
	"github.com/kurtosis-tech/stacktrace"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	grpcServerStopGracePeriod = 5 * time.Second
)

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
type KurtosisModuleExecutor struct {
	configurator KurtosisModuleConfigurator
}

func NewKurtosisModuleExecutor(configurator KurtosisModuleConfigurator) *KurtosisModuleExecutor {
	return &KurtosisModuleExecutor{configurator: configurator}
}

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
func (executor KurtosisModuleExecutor) Run() error {

	serializedCustomParams, err := getEnvVar(kurtosis_module_docker_api.SerializedCustomParamsEnvVar, "the serialized custom params that the module will consume")
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred when trying to get the serialized custom params environment variable")
	}

	module, err := executor.configurator.ParseParamsAndCreateExecutableModule(serializedCustomParams)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing the serialized custom params and creating the module")
	}

	apiContainerSocket, err := getEnvVar(kurtosis_module_docker_api.ApiContainerSocketEnvVar, "the socket value used in API container connection")
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
		kurtosis_module_docker_api.EnclaveDataDirMountpoint,
	)

	serviceImpl := newExecutableModuleServiceImpl(module, networkCtx)
	serviceImplRegistrationFunc := func(grpcServer *grpc.Server) {
		kurtosis_module_rpc_api_bindings.RegisterExecutableModuleServiceServer(grpcServer, serviceImpl)
	}

	grpcServer := grpc_server.NewMinimalGRPCServer(
		kurtosis_module_rpc_api_consts.ListenPort,
		kurtosis_module_rpc_api_consts.ListenProtocol,
		grpcServerStopGracePeriod,
		[]func(desc *grpc.Server){
			serviceImplRegistrationFunc,
		},
	)
	if err := grpcServer.Run(); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the module GRPC server")
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
		return "", stacktrace.NewError("The '%v' environment variable was defined, but is emptystring", envVarName)
	}

	return envVarValue, nil
}

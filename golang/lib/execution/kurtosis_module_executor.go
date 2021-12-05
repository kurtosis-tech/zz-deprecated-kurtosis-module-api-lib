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
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/module_launch_api"
	grpc_server "github.com/kurtosis-tech/minimal-grpc-server/golang/server"
	"github.com/kurtosis-tech/stacktrace"
	"google.golang.org/grpc"
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
	args, err := module_launch_api.GetArgsFromEnv()
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the module container args from the environment")
	}
	apiContainerSocket := args.ApiContainerSocket
	serializedCustomParams := args.SerializedCustomParams
	enclaveId := enclaves.EnclaveID(args.EnclaveID)
	listenPortNum := args.ListenPortNum
	enclaveDataDirMountpoint := args.EnclaveDataDirMountpoint;

	module, err := executor.configurator.ParseParamsAndCreateExecutableModule(serializedCustomParams)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating the module using serialized custom params '%v'", serializedCustomParams)
	}

	// TODO SECURITY: Use HTTPS to verify we're hitting the correct API container
	conn, err := grpc.Dial(apiContainerSocket, grpc.WithInsecure())
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred dialling the API container at '%v'", apiContainerSocket)
	}

	apiClient := kurtosis_core_rpc_api_bindings.NewApiContainerServiceClient(conn)
	enclaveCtx := enclaves.NewEnclaveContext(
		apiClient,
		enclaveId,
		enclaveDataDirMountpoint,
	)

	serviceImpl := newExecutableModuleServiceImpl(module, enclaveCtx)
	serviceImplRegistrationFunc := func(grpcServer *grpc.Server) {
		kurtosis_core_rpc_api_bindings.RegisterExecutableModuleServiceServer(grpcServer, serviceImpl)
	}

	grpcServer := grpc_server.NewMinimalGRPCServer(
		listenPortNum,
		grpcServerStopGracePeriod,
		[]func(desc *grpc.Server){
			serviceImplRegistrationFunc,
		},
	)
	if err := grpcServer.RunUntilInterrupted(); err != nil {
		return stacktrace.Propagate(err, "An error occurred running the module GRPC server")
	}

	return nil
}

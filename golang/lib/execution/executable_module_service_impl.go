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
	"context"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/enclaves"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/kurtosis_module_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-module-api-lib/golang/lib/kurtosis_modules"
	"github.com/kurtosis-tech/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
)

type executableModuleServiceImpl struct {
	// This embedding is required by gRPC
	kurtosis_module_rpc_api_bindings.UnimplementedExecutableModuleServiceServer
	module     kurtosis_modules.ExecutableKurtosisModule
	enclaveCtx *enclaves.EnclaveContext
}

func newExecutableModuleServiceImpl(
	module kurtosis_modules.ExecutableKurtosisModule,
	enclaveCtx *enclaves.EnclaveContext,
) *executableModuleServiceImpl {
	return &executableModuleServiceImpl{
		module:     module,
		enclaveCtx: enclaveCtx,
	}
}

func (server *executableModuleServiceImpl) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (server *executableModuleServiceImpl) Execute(ctx context.Context, args *kurtosis_module_rpc_api_bindings.ExecuteArgs) (*kurtosis_module_rpc_api_bindings.ExecuteResponse, error) {
	result, err := server.module.Execute(server.enclaveCtx, args.ParamsJson)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred executing the module")
	}
	executeResponse := &kurtosis_module_rpc_api_bindings.ExecuteResponse{
		ResponseJson: result,
	}

	return executeResponse, nil
}

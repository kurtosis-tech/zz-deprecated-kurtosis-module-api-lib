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
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/kurtosis-lambda"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KurtosisLambdaServiceServer struct {
	// This embedding is required by gRPC
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer
	kurtosisLambda kurtosis_lambda.KurtosisLambda
	networkCtx     *networks.NetworkContext
}

func NewKurtosisLambdaServiceServer(kurtosisLambda kurtosis_lambda.KurtosisLambda, networkCtx *networks.NetworkContext) *KurtosisLambdaServiceServer {
	return &KurtosisLambdaServiceServer{kurtosisLambda: kurtosisLambda, networkCtx: networkCtx}
}

func (lambdaService *KurtosisLambdaServiceServer) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (lambdaService *KurtosisLambdaServiceServer) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	result, err := lambdaService.kurtosisLambda.Execute(lambdaService.networkCtx, args.ParamsJson)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred executing the KurtosisLambda")
	}
	executeResponse := &kurtosis_lambda_rpc_api_bindings.ExecuteResponse{
		ResponseJson: result,
	}

	return executeResponse, nil
}

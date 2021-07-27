/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
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

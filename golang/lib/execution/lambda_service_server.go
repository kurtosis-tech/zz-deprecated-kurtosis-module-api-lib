/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
	"context"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/lambda"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LambdaServiceServer struct {
	// This embedding is required by gRPC
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer
	lambda lambda.Lambda
	client kurtosis_lambda_rpc_api_bindings.LambdaServiceClient
}

func NewLambdaServiceServer(lambda lambda.Lambda, client kurtosis_lambda_rpc_api_bindings.LambdaServiceClient) *LambdaServiceServer {
	return &LambdaServiceServer{lambda: lambda, client: client}
}

func (lambdaService *LambdaServiceServer) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	panic("implement me")
}

func (lambdaService *LambdaServiceServer) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	panic("implement me")
}


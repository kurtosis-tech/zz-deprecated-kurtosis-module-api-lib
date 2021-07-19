/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
	"context"
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/lambda"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
	"reflect"
)

type LambdaServiceServer struct {
	// This embedding is required by gRPC
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer
	lambda     lambda.Lambda
	networkCtx *networks.NetworkContext
}

func NewLambdaServiceServer(lambda lambda.Lambda, networkCtx *networks.NetworkContext) *LambdaServiceServer {
	return &LambdaServiceServer{lambda: lambda, networkCtx: networkCtx}
}

func (lambdaService *LambdaServiceServer) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {

	err := lambdaService.lambda.IsAvailable()
	if err != nil {
		return &emptypb.Empty{}, stacktrace.Propagate(err, "Lambda %v is not available", reflect.TypeOf(lambdaService.lambda).String())
	}

	return &emptypb.Empty{}, nil
}

func (lambdaService *LambdaServiceServer) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	panic("implement me")
}

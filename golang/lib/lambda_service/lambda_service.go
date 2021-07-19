/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package lambda_service

import (
	"context"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_rpc_api_bindings"
	"github.com/palantir/stacktrace"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LambdaService struct {
	client kurtosis_lambda_rpc_api_bindings.LambdaServiceClient
}

func NewLambdaService(client kurtosis_lambda_rpc_api_bindings.LambdaServiceClient) *LambdaService {
	return &LambdaService{client: client}
}

func (lambda *LambdaService) IsAvailable() (bool, error) {
	_, err := lambda.client.IsAvailable(context.Background(), &emptypb.Empty{})
	if err != nil {
		return false, stacktrace.Propagate(err, "An error occurred when trying to check Lambda service availability")
	}
	return true, nil
}

func (lambda *LambdaService) Execute(argsJsonStr string) (responseJsonStr string, resultErr error) {
	args := &kurtosis_lambda_rpc_api_bindings.ExecuteArgs{
		ParamsJson: argsJsonStr,
	}

	resp, err := lambda.client.Execute(context.Background(), args)
	if err != nil {
		return "", stacktrace.Propagate(err, "An error occurred executing Lambda")
	}
	return resp.ResponseJson, nil
}

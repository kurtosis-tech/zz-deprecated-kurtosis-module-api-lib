/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/lambda"
)

type LambdaConfigurator interface {

	SetLogLevel(logLevelStr string) error

	ParseParamsAndCreateLambda(paramsJsonStr string) (lambda.Lambda, error)
}

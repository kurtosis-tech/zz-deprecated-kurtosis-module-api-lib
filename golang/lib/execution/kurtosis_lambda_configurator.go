/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution

import (
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/kurtosis-lambda"
)

type KurtosisLambdaConfigurator interface {
	ParseParamsAndCreateKurtosisLambda(serializedCustomParamsStr string) (kurtosis_lambda.KurtosisLambda, error)
}

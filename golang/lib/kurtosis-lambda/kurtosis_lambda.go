/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package kurtosis_lambda

import "github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"

type KurtosisLambda interface {
	Execute(networkCtx *networks.NetworkContext, serializedParams string) (serializedResult string, resultError error)
}

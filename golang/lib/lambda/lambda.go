/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package lambda

import "github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"

type Lambda interface {

	IsAvailable() error

	Execute(networkCtx *networks.NetworkContext, serializedParams string) (serializedResult string, resultError error)
}

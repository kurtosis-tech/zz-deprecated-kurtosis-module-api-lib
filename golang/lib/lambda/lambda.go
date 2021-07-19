/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package lambda

type Lambda interface {

	IsAvailable() bool

	Execute(serializedParams string) (serializedResult string, resultError error)
}

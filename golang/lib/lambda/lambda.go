/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package lambda

type Lambda interface {

	IsAvailable() error

	Execute(serializedParams string) (serializedResult string, resultError error)
}

/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package impl

type ExampleLambda struct {
	isKurtosisCoreDevMode bool
}

func NewExampleLambda(isKurtosisCoreDevMode bool) *ExampleLambda {
	return &ExampleLambda{isKurtosisCoreDevMode: isKurtosisCoreDevMode}
}

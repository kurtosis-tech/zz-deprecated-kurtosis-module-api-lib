/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package example

type ExampleLambda struct {
	isKurtosisCoreDevMode bool
}

func NewExampleLambda(isKurtosisCoreDevMode bool) *ExampleLambda {
	return &ExampleLambda{isKurtosisCoreDevMode: isKurtosisCoreDevMode}
}

func (e ExampleLambda) IsAvailable() error {
	return nil
}

func (e ExampleLambda) Execute(serializedParams string) (serializedResult string, resultError error) {
	panic("implement me")
}

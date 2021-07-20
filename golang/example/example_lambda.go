/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package example

import (
	"github.com/kurtosis-tech/kurtosis-client/golang/lib/networks"
	"github.com/sirupsen/logrus"
)

type ExampleLambda struct {
}

func NewExampleLambda() *ExampleLambda {
	return &ExampleLambda{}
}

func (e ExampleLambda) IsAvailable() error {
	return nil
}

func (e ExampleLambda) Execute(networkCtx *networks.NetworkContext, serializedParams string) (serializedResult string, resultError error) {
	logrus.Infof("Example Lambda receives serializedParams '%v'", serializedParams)
	serializedResult = "Example Lambda Module successful execution"
	logrus.Info("Example Lambda executed successfully")
	return serializedResult, nil
}

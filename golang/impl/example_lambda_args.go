/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package impl

type ExampleLambdaArgs struct {
	// Indicates that this testsuite is being run as part of CI testing in Kurtosis Core
	IsKurtosisCoreDevMode bool		`json:"isKurtosisCoreDevMode"`
}

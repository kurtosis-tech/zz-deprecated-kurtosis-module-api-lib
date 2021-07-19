/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package example

type ExampleLambdaArgs struct {
	// Indicates that this Lambda is being run as part of CI testing in Kurtosis Core
	IsKurtosisCoreDevMode bool `json:"isKurtosisCoreDevMode"`
	// Indicates the log level for this Lambda implementation
	logLevel string `json:"logLevel"`
}

package main

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/example"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/kurtosis_lambda_docker_api"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/execution"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

func main() {

	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<
	configurator := example.NewExampleLambdaConfigurator()
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<

	apiContainerSocketArg, serializedCustomParamsArg, err := getArgs()
	if err != nil {
		logrus.Errorf("An error occurred getting the Lambda args:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}

	lambdaExecutor := execution.NewLambdaExecutor(apiContainerSocketArg, serializedCustomParamsArg, configurator)
	if err := lambdaExecutor.Run(); err != nil {
		logrus.Errorf("An error occurred running the lambda executor:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}

func getArgs() (string, string, error) {
	apiContainerSocketEnvVar, found := os.LookupEnv(kurtosis_lambda_docker_api.ApiContainerSocketEnvVar)
	if !found {
		return "", "", stacktrace.NewError("No API container socket environment variable '%v' defined", kurtosis_lambda_docker_api.ApiContainerSocketEnvVar)
	}

	serializedCustomParamsEnvVar, found := os.LookupEnv(kurtosis_lambda_docker_api.SerializedCustomParamsEnvVar)
	if !found {
		return "", "", stacktrace.NewError("No serialized custom params environment variable '%v' defined", kurtosis_lambda_docker_api.SerializedCustomParamsEnvVar)
	}

	return apiContainerSocketEnvVar, serializedCustomParamsEnvVar, nil
}

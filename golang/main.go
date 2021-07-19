package golang

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/impl"
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
	configurator := impl.NewExampleLambdaConfigurator()
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<

	lambdaExecutor := execution.NewLambdaExecutor(*kurtosisLambdaApiSocketArg, *logLevelArg, *customParamsJsonArg, configurator)
	if err := lambdaExecutor.Run(); err != nil {
		logrus.Errorf("An error occurred running the lambda executor:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}

//TODO refactor this, create a Configuration Struct
func getRunConfigurations() (string, string, string, error) {
	apiContainerSocketEnvVar, found := os.LookupEnv(kurtosis_lambda_docker_api.ApiContainerSocketEnvVar)
	if !found {
		return "", "", "", stacktrace.NewError("No API container socket environment variable '%v' defined", kurtosis_lambda_docker_api.ApiContainerSocketEnvVar)
	}

}

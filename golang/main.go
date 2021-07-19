package golang

import (
	"flag"
	"fmt"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/impl"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/execution"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

func main() {

	customParamsJsonArg := flag.String(
		"custom-params-json",
		"{}",
		"JSON string containing custom data that the lambda will deserialize to modify runtime behaviour",
	)

	kurtosisLambdaApiSocketArg := flag.String(
		"kurtosis-lambda-api-socket",
		"",
		"Socket in the form of address:port of the Kurtosis Lambda API",
	)

	logLevelArg := flag.String(
		"log-level",
		"",
		"String indicating the loglevel that the lambda should output with",
	)

	flag.Parse()

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

package impl

import (
	"encoding/json"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/lambda"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

type ExampleLambdaConfigurator struct{}

func NewExampleLambdaConfigurator() *ExampleLambdaConfigurator {
	return &ExampleLambdaConfigurator{}
}

func (t ExampleLambdaConfigurator) SetLogLevel(logLevelStr string) error {
	level, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing loglevel string '%v'", logLevelStr)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	return nil
}

func (t ExampleLambdaConfigurator) ParseParamsAndCreateLambda(paramsJsonStr string) (lambda.Lambda, error) {
	paramsJsonBytes := []byte(paramsJsonStr)
	var args ExampleLambdaArgs
	if err := json.Unmarshal(paramsJsonBytes, &args); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred deserializing the Lambda params JSON")
	}

	lambda := NewExampleLambda(args.IsKurtosisCoreDevMode)

	return lambda, nil
}

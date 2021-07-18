package kurtosis_lambda_docker_api

const (
	// Name of Docker environment variable where JSON-serialized Kurtosis params will be sent
	KurtosisParamsJsonEnvVar = "KURTOSIS_PARAMS_JSON"

	// Name of Docker environment variable where users can send custom params for modifying how the lambda behaves
	CustomParamsJsonEnvVar = "CUSTOM_PARAMS_JSON"

	// Location on the Lambda Docker container where the Kurtosis volume will be mounted
	ExecutionVolumeMountpoint = "/kurtosis-execution-volume"
)



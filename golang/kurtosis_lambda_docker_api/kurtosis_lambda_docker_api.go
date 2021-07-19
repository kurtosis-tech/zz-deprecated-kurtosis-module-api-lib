package kurtosis_lambda_docker_api

const (
	// A string defining the loglevel that the Lambda should output
	// This string should be meaningful to the user, since Kurtosis won't know what logging framework they use
	LogLevelEnvVar = "LOG_LEVEL"

	// IP:port of the Kurtosis API container
	ApiContainerSocketEnvVar = "API_CONTAINER_SOCKET"

	// Arbitrary serialized data that the Lambda can consume at startup to modify its behaviour
	// Analogous to the "constructor"
	SerializedCustomParamsEnvVar = "SERIALIZED_CUSTOM_PARAMS"

	// Location on the Lambda Docker container where the Kurtosis volume will be mounted
	ExecutionVolumeMountpoint = "/kurtosis-execution-volume"
)

// IP:port of the Kurtosis API container
export const API_CONTAINER_SOCKET_ENV_VAR: string = "API_CONTAINER_SOCKET";

// Arbitrary serialized data that the module can consume at startup to modify its behaviour
// Analogous to the "constructor"
export const SERIALIZED_CUSTOM_PARAMS_ENV_VAR: string = "SERIALIZED_CUSTOM_PARAMS";

// Location on the module Docker container where the Kurtosis volume will be mounted
export const EXECUTION_VOLUME_MOUNTPOINT: string = "/kurtosis-execution-volume";

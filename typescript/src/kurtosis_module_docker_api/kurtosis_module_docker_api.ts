// IP:port of the Kurtosis API container
export const API_CONTAINER_SOCKET_ENV_VAR: string = "API_CONTAINER_SOCKET";

// Arbitrary serialized data that the module can consume at startup to modify its behaviour
// Analogous to the "constructor"
export const SERIALIZED_CUSTOM_PARAMS_ENV_VAR: string = "SERIALIZED_CUSTOM_PARAMS";

// Location on the module Docker container where the Kurtosis enclave data directory will be bind-mounted
export const ENCLAVE_DATA_DIR_MOUNTPOINT: string = "/kurtosis-enclave-data"
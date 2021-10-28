/*
 *    Copyright 2021 Kurtosis Technologies Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package kurtosis_module_docker_api

const (
	// IP:port of the Kurtosis engine server
	EngineSocketEnvVar = "ENGINE_SOCKET"

	// ID of the enclave that the module is executing inside
	EnclaveIDEnvVar = "ENCLAVE_ID"

	// Arbitrary serialized data that the module can consume at startup to modify its behaviour
	// Analogous to the "constructor"
	SerializedCustomParamsEnvVar = "SERIALIZED_CUSTOM_PARAMS"

	// Location on the module Docker container where the Kurtosis enclave data directory will be bind-mounted
	EnclaveDataDirMountpoint = "/kurtosis-enclave-data"
)

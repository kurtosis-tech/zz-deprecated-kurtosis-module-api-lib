# TBD

# 0.23.0

### Changes
* Uses `kurtosis-sdk` instead of `kurtoiss-core-sdk`

### Breaking Changes
* Uses `kurtosis-sdk` instead of `kurtoiss-core-sdk`

# 0.22.5

### Changes
* Use `kurtosis-core-sdk` instead of `kurtosis-core-api-lib`

# 0.22.4
### Changes
* Upgraded core to 1.59.5

# 0.22.3
### Changes
* Upgraded core to 1.59.3

# 0.22.2
### Changes
* Upgraded core to 1.59.2


# 0.22.1
### Changes
* Upgraded core to 1.59.1 and engine to 1.34.0

# 0.22.0
### Changes
* Upgraded core to 1.59.0

### Breaking changes
* Upgraded core to 1.59.0
  * Added `privateIPAddrPlaceholder` field to the `ContainerConfig object`. This defaults to `"KURTOSIS_IP_ADDR_PLACEHOLDER"` in case the user doesn't supply it. Kurtosis will replace the placeholder in the entry point args, env vars & cmd args with the private ip address of the container while starting the container.
  * Changed the `EnclaveContext.AddService`, `EnclaveContext.AddServices`, `EnclaveContext.AddServiceToPartition` and `EnclaveContext.AddServicesToPartition` to accept a `ContainerConfig` object instead of a `ContainerConfigSupplier`
  * Now the user can simply pass in a `containerConfig ContainerConfig` over supplying a `serviceConfigSupplier func(ipAddr string) (*services.ContainerConfig, error)` which supports a `privateIPAddrPlaceholder` mentioned above.


# 0.21.2
### Fixes
* Upgraded core to 1.58.3 which fixes how we handle json in the `RenderTemplatesToFilesArtifact` end point.

### Changes
* Upgraded core to 1.58.2 which exposes a `EnclaveContext.RenderTemplates` method in the SDK.

# 0.21.1
### Changes
* Upgrade to core to 1.58.1

# 0.21.0
### Changes
* Upgrade Core to 1.58.0 which changes the behaviour of `EnclaveContext.UploadFiles`

### Breaking Changes
* Upgraded Core to 1.58.0 - when you now use the `EnclaveContext.UploadFiles` for a directory, the resulting archive artifact will contain the _contents_ of the directory in the root of the archive. The previous implementation would put the _contents_ inside a folder with the same name as the directory, the new implementation avoids this nesting.
  * User who access the artifact on their service should now directly access the files and directories(within the directory uploaded) at the mount point. Any previous paths that contain references to the name of the directory uploaded should be replaced to not have that name. The directory structure within the directory that was uploaded has been preserved.

# 0.20.0
### Breaking Changes
* Upgrade to Core 1.57.6
  * Users should restart their Kurtosis engine restart their Kurtosis engine
  
# 0.19.0
### Breaking Changes
* Upgrade to Core 1.57.3 and engine to 1.31.0
  * Users should restart their Kurtosis engine restart their Kurtosis engine

# 0.18.1

# 0.18.0
### Breaking Changes
* Upgrade to Core 1.57.0
  * Users should restart their Kurtosis engine restart their Kurtosis engine

### Changes
* Migrate repo to use internal cli tool `kudet` for releasing
* Merge `develop` into `master`


# 0.17.0
### Breaking Changes
* Upgraded to Core 1.55.2

# 0.16.0
### Breaking Changes
* Upgraded to Core 1.54.1

# 0.15.0
### Breaking Changes
* Upgraded to Core 1.46.0, to work with latest version of Kurtosis CLI
    * Users should upgrade to the latest engine and run `kurtosis engine restart`

# 0.14.1
### Changes
* Use Core 1.45.3 (API breaks were actually intended to go in last version)

# 0.14.0
### Breaking Changes
* Upgraded engine API lib to 1.17.3
    * Users should see [the changes done in v1.44.0 and 1.45.0, and follow the remediation steps there](https://docs.kurtosistech.com/kurtosis/historical-core-changelog)

# 0.13.0
### Features
* Upgrade to Kurt Core 1.41.1

### Breaking Changes
* Upgraded Kurt Core API lib to 1.41.1

# 0.12.3
### Fixes
* Upgrade to core-api-lib 1.36.8, to try and fix an issue where the Protobuf "empty" type isn't getting propagated

# 0.12.2
### Changes
* Switch to using `@grpc/grpc-js` as the `grpc` package is deprecated, upgrading the following dependencies to support this:
    * `kurtosis-core-api-lib` -> 1.36.7
    * `minimal-grpc-server` -> 0.6.0

# 0.12.1
### Fixes
* The `package.json` now correctly declares a requirement on Node >= 16.13.0

# 0.12.0
### Features
* Upgraded to `minimal-grpc-server` 0.5.0
* Added a root `scripts/build.sh` to build all languages
* Added test to ensure that a new Kurt Core version with an API break will remind the user to add a breaking change in this library's changelog

### Changes
* Replaced the old `kurtosis-client` with `kurtosis-core-api-lib`

### Fixes
* A `nil` value passed to `stacktrace.Propagate` now panics

### Removals
* Removed the Protobuf API file & generated binding files, as the module's API is now defined in [Kurtosis Core API Lib](https://github.com/kurtosis-tech/kurtosis-core-api-lib) rather than here
* Removed the Docker & RPC API constants packages, as the information they used to contain now comes from the args that the module container is passed in

### Breaking Changes
* The `kurtosis-client` dependency has been replaced with the [Kurtosis Core API Lib](https://github.com/kurtosis-tech/kurtosis-core-api-lib) dependency
    * Users of any module languages (Go or Typescript) will need to:
        * Replace all instances of the old `NetworkContext` with the new `EnclaveContext`
    * Go module users will need to:
        * Replace the `github.com/kurtosis-tech/kurtosis-client` dependency in their `go.mod` file with a dependency on `github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang`
        * In all `import` statements, replace all instances of `github.com/kurtosis-tech/kurtosis-client` with `github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang`
    * Typescript module users will need to:
        * Replace the dependency on `kurtosis-clietn` in their `package.json` with a dependency on `kurtosis-core-api-lib`

# 0.11.1
### Changes
* Upgraded to Kurt Client 0.20.0, which is compatible with new bind-mount style enclave data dirs

# 0.11.0
### Changes
* The execution volume mountpoint constants have been renamed to reflect the new bind-mounting strategy

### Breaking Changes
* Renamed the `ExecutionVolumeMountpoint` constant to `EnclaveDataDirMountpoint` to reflect the new bind-mounting strategy
    * Users should rename their Golang & Typescript constants appropriately

# 0.10.0
### Features
* Added documentation for the `KurtosisModuleExecutor` class

### Fixes
* Upgrade to [Kurtosis Client 0.19.0](https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md#0190), which fixes Typescript `SharedPath.GetChildPath` being accidentally uppercased

### Changes
* All "Lambda" instances in the API have been renamed to "module"

### Removals
* Actually remove the Lambda registry Markdown doc, which should have been done in 0.9.0

### Breaking Changes
* This repo's Go module has been renamed to `github.com/kurtosis-tech/kurtosis-module-api-lib/golang`
    * Uses should update their imports accordingly
* This repo's NPM module has been renamed to `kurtosis-module-api-lib`
* Typescript `SharedPath.GetChildPath` renamed to `SharedPath.getChildPath`
* The Kurtosis Client API has received significant updates from "Lambda" to "module"
    * Users should follow the remediation steps outlined [here](https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md#0190)
* Several classes & functions were renamed, and users should modify their code appropriately:
    * `KurtosisLambda` was renamed to `ExecutableKurtosisModule`
    * `KurtosisLambdaExecutor` was renamed to `KurtosisModuleExecutor`
    * `KurtosisLambdaConfigurator` was renamed to `KurtosisModuleConfigurator`
        * The `parseParamsAndCreateLambda` function was renamed `parseParamsAndCreateExecutableModule`

# 0.9.2
### Features
* Upgraded to [Kurtosis Client 0.17.3](https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md#0173)

# 0.9.1
### Features
* Upgraded Kurt client dependency to the latest version [Kurt Client API 0.17.1](https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md#0171)

### Removals
* Removed Lambda registry (now lives in the official docs)

# 0.9.0
### Changes
* Add [Ethereum Kurtosis Lambda](https://github.com/kurtosis-tech/ethereum-kurtosis-lambda/) in Kurtosis Lambda registry
* Upgraded to `kurtosis-client` 0.16.0, which returns log output strings rather than bytes from `ServiceContext.execCommand`

### Breaking Changes
* `ServiceContext.execCommand` returns strings rather than bytes
    * Users should switch to using the string directly, without decoding

# 0.8.0
### Fixes
* Upgraded to Kurt Client 0.15.0 to fix a typo in ContainerRunConfigBuilder

### Breaking Changes
* Upgrade to Kurt Client 0.15.0 (see break remediation [here](https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md))

# 0.7.8
### Changes
* Upgraded Kurtosis Client from 0.13.2 -> 0.13.7

# 0.7.7
### Fixes
* Fixed a bug with return `null` instead of a `google_protobuf_empty_pb.Empty` object for Typescript `isAvailable` method

# 0.7.6
### Fixes
* Use minimal-grpc-server 0.3.5, which fixes a bug with the bind URL

# 0.7.5
### Fixes
* Use minimal-grpc-server 0.3.4, which fixes a bug with starting the server

# 0.7.4
### Fixes
* Use minimal-grpc-server 0.3.3, which unpins the Node engine version from `14.17.1`

# 0.7.3
### Changes
* Upgraded to Kurtosis Client 0.13.2, which contains a bunch of bugfixes surfaced by strict mode

# 0.7.2
### Features
* Turn on Typescript strict mode, for safer code

### Fixes
* Several bugs resulting from strict mode being enabled

# 0.7.1
### Fixes
* Unpin Typescript library Node engine version to be `>=14.17.0`
* Typescript library was missing an `index.ts`

# 0.7.0
### Changes
* Correct for `minimal-grpc-server` module move
* Switch to docs-checker orb
* Use the devtools version of the package-updating script

### Features
* Set up Typescript build infra
* Add Typescript implementation

### Fixes
* Fixed hyphenated (rather than underscored) directory name in the Golang module, to follow Go naming conventions
* Added docs links to all classes

### Breaking Changes
* Renamed `github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/kurtosis-lambda` package -> `github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/kurtosis_lambda`
    * Users should rename the package

# 0.6.0
### Changes
* Upgraded Kurtosis Client from 0.10.0 -> 0.11.0
* Added Apache-2 license

### Features
* Added a known Lambdas registry

### Breaking Changes
* Upgraded Kurtosis Client from 0.10.0 -> 0.11.0
    * Users should follow the remediation steps in https://github.com/kurtosis-tech/kurtosis-client/blob/develop/docs/changelog.md#0110 

# 0.5.1
### Changes
* Updated Kurtosis Client dependency version
* Updated `NewNetworkContext` arguments number when it's called in `KurtosisLambdaExecutor.Run` method

# 0.5.0

### Features
* Add library's documentation

### Breaking Changes
* Renamed `Lamba` -> `KurtosisLambda`
* Renamed `LambdaConfigurator` -> `KurtosisLambdaConfigurator`
* Renamed `LambdaExecutor` -> `KurtosisLambdaExecutor`
* Renamed `LambdaServiceServer` -> `KurtosisLambdaServiceServer`

# 0.4.1
### Features
* Add `LambdaConfigurator` interface; users should create their own custom implementation of this to configure their own Lambda
* Add `LambdaExecutor` which accepts a `LambdaConfigurator` implementation and is responsible for starting the Lambda server
* Add `LambdaServiceServer` which is the RPC server implementation of `kurtosis-lambda-rpc-api`
* Add `Lambda` interface; users should create their own custom Lambda implementations based on this contract which will be returned by the `LambdaConfigurator`
* Added much more detail to the README

# 0.4.0
### Removed
* Removed the Lambda loglevel environment variable

### Breaking Changes
* Removed the Lambda loglevel environment variable as a firstclass concept, since that should be handled by the Lambda itself 
    * Users depending on this variable should push loglevel-setting into the Lambda custom initialization params

# 0.3.1
### Features
* Add the set of Docker envvars as constants that can be used by both the API container (to send the constants) and Lambda (to receive the constants)
* Added a Lambda custom params Docker environment variable, analogous to a constructor, so Lambdas can accept data upon creation that will modify their behaviour

# 0.3.0
### Changes
* Renamed library name

### Breaking Changes
* Removed args & response objects to `IsAvailable` endpoint

# 0.2.0
### Breaking Changes
* Renamed library/module name to `kurtosis-lambda-api-lib`

### Breaking Changes
* Renamed Go module name to `kurtosis-lambda-api-lib`
    * Users will need to replace the old module name, `kurtosis-lambda-client`, with `kurtosis-lambida-api-lib`

# 0.1.3
* Add `kurtosis_lambda_rpc_api_consts` which contains protocol and port number
* Add an explicit `kurtosis_` in the API binding packages, to clarify that they're Kurtosis specific

# 0.1.2
### Features
* Defined Lambda Service GRPC contract
* Add `regenerate-protobuf-output` script file which uses the latest version of `generate-protobuf-bindings` from the devtools repo to generate bindings files
* Generated the Golang and Typescript bindings files
* Add the documentation of the library

# 0.1.1
* Init commit
* Add base directories and files
  * Add changelog file
  * Add golang basic structure
  * Add release script
* Set up the Circle-CI workflow

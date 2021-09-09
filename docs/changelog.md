# TBD
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

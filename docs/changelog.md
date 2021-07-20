# TBD
### Features
* Add `lambda_configurator` interface; users should create custom implementation of this to configure their own Lambda
* Add `lambda_executor` which is responsible for starts the server, using an implementation of `lambda_configurator`, with the Lambda implementation
* Add `lambda_service_service` which is the RPC server implementation of `kurtosis-lambda-rpc-api`
* Add `lambda` interface; users should create their own custom Lambda implementations based on this contract

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

# 0.1.1
* Init commit
* Add base directories and files
  * Add changelog file
  * Add golang basic structure
  * Add release script
* Set up the Circle-CI workflow

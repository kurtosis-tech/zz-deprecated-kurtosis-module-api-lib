# TBD

### Features
* Defined Lambda Service GRPC contract
* Add `regenerate-protobuf-output` script file which uses the latest version of `generate-protobuf-bindings` from the devtools repo to generate bindings files
* Generated the Golang and Typescript bindings files
* Add `kurtosis_lambda_rpc_api_consts` which contains protocol and port number
* Add an explicit `kurtosis_` in the API binding packages, to clarify that they're Kurtosis specific

# 0.1.1
* Init commit
* Add base directories and files
  * Add changelog file
  * Add golang basic structure
  * Add release script
* Set up the Circle-CI workflow

Kurtosis Lambda API LIB 
=======================

### What's a Lambda?
A Kurtosis Lambda is a modular chunk of code with a connection to [the Kurtosis engine](https://docs.kurtosistech.com/) that can be run inside a testnet and distributed to others. This allows for tests to be written in seconds from preexisting components.

Some examples of Lambda usage:

- Smart contract tests on top of an Ethereum Lambda that handles all the ETH network setup for the dApp developers
- Correctness tests that incorporate a fuzz-testing Lambda to spam garbage data at your endpoints
- Failure tests using a Lambda that randomly restarts services in your network, a la Netflix' Chaos Monkey
- Partition tolerance tests using a Lambda that periodically repartitions & heals your network

More formally, a Lambda is:

- A gRPC server
- With a connection to the Kurtosis engine
- And a single endpoint for executing the Lambda
- Written in an arbitrary language
- That is packaged inside a Docker image

### What's in this repo?

- A definition of the Kurtosis Lambda API
- A client for communicating with that API
- Scaffolding for building a Lambda
- [A list of known Lambdas](./docs/lambda-registry.md)

### How do I write my own Lambda?
For right now, an example Lambda is provided [here](https://github.com/kurtosis-tech/datastore-army-lambda). A more detailed tutorial will be available soon.

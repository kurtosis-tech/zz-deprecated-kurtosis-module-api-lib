Kurtosis Lambda API Lib Documentation
=====================================
This documentation describes all the necessary objects to configure, create and run a custom implementation of a Kurtosis Lambda.


Kurtosis Lambda
---------------
This interface defines the behavior of a Kurtosis Lambda which is a way to package and make available preconfigured services that users can easily load and use in their Kurtosis tests.

### Execute([NetworkContext][networkcontext] networkContext, String serializedParams) -\> (String serializedResult)
This function executes the only functionality that contains an implementation of a Kurtosis Lambda.
It receives an unlimited amount of serialized parameters and also returns an unlimited amount of serialized values. The implementation of this method must be in charge of validating and sanitizing the received and returned values.

**Args**

* `networkContext`: The representation of the network which must be used to interact with the API core, and the others services running in the network
* `serializedParams`: Serialized parameter data that will be passed to the Kurtosis Lambda to control overall the behaviour.

**Returns**

* `serializedResult`: Serialized data containing the results of executing the Kurtosis Lambda function.


KurtosisLambdaConfigurator
--------------------------
Object responsible for create and configure a custom implementation of a Kurtosis Lambda.

### ParseParamsAndCreateKurtosisLambda(String serializedCustomParamsStr) ([KurtosisLambda][kurtosislambda] kurtosisLambda)
Creates and configure a Kurtosis Lambda defined by the received params which previously were parsed

**Args**

* `serializedCustomParamsStr`: Serialized data containing params to create and configure an implementation of a Kurtosis Lambda. Each implementation must define amount and type of the params.

**Returns**

* `kurtosisLambda`: The object that represents the implementation of a Kurtosis Lambda 

KurtosisLambdaExecutor
----------------------

### Run ()
Run the Kurtosis Lambda RPC server and connect with API container


KurtosisLambdaServiceServer
---------------------------
#TODO: Should we describe this, because users must not directly use this object 

NetworkContext
--------------
This Kurtosis-provided class is the lowest-level representation of a test network, and provides methods for inspecting and manipulating the network.

[networkcontext]: #networkcontext
[kurtosislambda]: #kurtosis-lambda
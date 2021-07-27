Kurtosis Lambda API Lib Documentation
=====================================
This documentation describes all the necessary objects to configure, create and run a custom implementation of a Kurtosis Lambda.


Kurtosis Lambda
---------------
This interface defines the behavior of a Kurtosis Lambda which is a way to package and make available preconfigured services that users can easily load and use in their Kurtosis tests.

### execute([NetworkContext][networkcontext] networkContext, String serializedParams) -\> (String serializedResult)
This function executes the functionality contained inside this implementation of Kurtosis Lambda.
It receives parameters, serialized in an implementation-specific format, and returns a response, also serialized in an implementation-specific format. The implementation of this method is in charge of deserializing, validating, and sanitizing the received parameters.

**Args**

* `networkContext`: The representation of the network, which is used to manipulate the network.
* `serializedParams`: Serialized parameter data to control the behaviour of the Kurtosis Lambda.

**Returns**

* `serializedResult`: Serialized data containing the results of executing the Kurtosis Lambda function.


KurtosisLambdaConfigurator
--------------------------
Object responsible for creating and configuring a custom implementation of a Kurtosis Lambda.

### parseParamsAndCreateKurtosisLambda(String serializedCustomParamsStr) ([KurtosisLambda][kurtosislambda] kurtosisLambda)
Creates and configures a Kurtosis Lambda configured with the given serialized custom params.

**Args**

* `serializedCustomParamsStr`: Serialized data containing params to create and configure an implementation of a Kurtosis Lambda.

**Returns**

* `kurtosisLambda`: The object that represents the implementation of a Kurtosis Lambda


[networkcontext]: ../kurtosis-client/lib-documentation#networkcontext
[kurtosislambda]: #kurtosis-lambda
Kurtosis Module API Lib Documentation
=====================================
This documentation describes all the necessary objects to configure, create and run a custom Kurtosis module.


ExecutableKurtosisModule
------------------------
This interface defines the behavior of a Kurtosis module that responds to an "execute" command. It is intended for simple modules that only have one function.

### execute([EnclaveContext][enclavecontext] enclaveContext, String serializedParams) -\> (String serializedResult)
This function executes the logic contained inside this module. It receives parameters, serialized in an implementation-specific format, and returns a response, also serialized in an implementation-specific format. The implementation of this method is in charge of deserializing, validating, and sanitizing the received parameters.

**Args**

* `enclaveContext`: The representation of the Kurtosis enclave, which is used for adding services, removing services, etc.
* `serializedParams`: Serialized parameter data to control the behaviour of the Kurtosis module.

**Returns**

* `serializedResult`: Serialized data containing the results of running the module's execute function.


KurtosisModuleConfigurator
--------------------------
Object responsible for creating and configuring a custom executable Kurtosis module.

### parseParamsAndCreateExecutableModule(String serializedCustomParamsStr) -\> ([ExecutableKurtosisModule][executablekurtosismodule] module)
Creates and configures a Kurtosis module with an execute command, configured with the given serialized custom params.

**Args**

* `serializedCustomParamsStr`: Serialized data containing params to create and configure an implementation of the [ExecutableKurtosisModule][executablekurtosismodule] interface.

**Returns**

* `module`: An implementation of the [ExecutableKurtosisModule][executablekurtosismodule] interface representing a module that responds to an execute command.

KurtosisModuleExecutor
----------------------
Executor which accepts a [KurtosisModuleConfigurator][kurtosismoduleconfigurator] as part of its constructor to create an [ExecutableKurtosisModule][executablekurtosismodule] and serve it over a gRPC server.

### run()
Runs the gRPC server exposing the functionality of the Kurtosis module created by the [KurtosisModuleConfigurator][kurtosismoduleconfigurator] passed in at construction time.

[enclavecontext]: ../kurtosis/core-lib-documentation#enclavecontext
[kurtosismoduleconfigurator]: #kurtosismoduleconfigurator
[executablekurtosismodule]: #executablekurtosismodule

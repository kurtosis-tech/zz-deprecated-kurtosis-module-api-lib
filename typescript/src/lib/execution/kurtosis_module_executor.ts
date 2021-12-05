import { ApiContainerServiceClient } from "kurtosis-core-api-lib";
import { err, ok, Result } from "neverthrow";
import { API_CONTAINER_SOCKET_ENV_VAR, ENCLAVE_DATA_DIR_MOUNTPOINT, SERIALIZED_CUSTOM_PARAMS_ENV_VAR } from "../../kurtosis_module_docker_api/kurtosis_module_docker_api";
import { ExecutableKurtosisModule } from "../kurtosis_modules/executable_kurtosis_modules";
import { KurtosisModuleConfigurator } from "./kurtosis_module_configurator";
import * as grpc from 'grpc';
import { MinimalGRPCServer, TypedServerOverride } from "minimal-grpc-server";
import { ExecutableModuleServiceImpl } from "./executable_module_service_impl";

const GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS: number = 5;

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
export class KurtosisModuleExecutor {
    private readonly configurator: KurtosisModuleConfigurator;

    constructor(configurator: KurtosisModuleConfigurator) {
        this.configurator = configurator;
    }

    // Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
    public async run(): Promise<Result<null, Error>> {
        const getSerializedCustomParamsResult: Result<string, Error> = KurtosisModuleExecutor.getEnvVar(SERIALIZED_CUSTOM_PARAMS_ENV_VAR, "the serialized custom params that the module will consume");
        if (getSerializedCustomParamsResult.isErr()) {
            return err(getSerializedCustomParamsResult.error);
        }
        const serializedCustomParams: string = getSerializedCustomParamsResult.value;

        const createModuleResult: Result<ExecutableKurtosisModule, Error> = this.configurator.parseParamsAndCreateExecutableModule(serializedCustomParams);
        if (createModuleResult.isErr()) {
            return err(createModuleResult.error);
        }
        const module: ExecutableKurtosisModule = createModuleResult.value;

        const getApiContainerSocketResult: Result<string, Error> = KurtosisModuleExecutor.getEnvVar(API_CONTAINER_SOCKET_ENV_VAR, "the socket value used in API container connection");
        if (getApiContainerSocketResult.isErr()) {
            return err(getApiContainerSocketResult.error);
        }
        const apiContainerSocket: string = getApiContainerSocketResult.value;

        const apiClient: ApiContainerServiceClient = new ApiContainerServiceClient(apiContainerSocket, grpc.credentials.createInsecure());
        const networkCtx: NetworkContext = new NetworkContext(
            apiClient,
            ENCLAVE_DATA_DIR_MOUNTPOINT,
        );

        const serviceImpl: ExecutableModuleServiceImpl = new ExecutableModuleServiceImpl(
            module,
            networkCtx
        );
        const serviceImplRegistrationFunc: { (server: TypedServerOverride): void; }[] = [
            (server: TypedServerOverride) => {
                server.addTypedService(ExecutableModuleServiceService, serviceImpl);
            }
        ];

        const grpcServer: MinimalGRPCServer = new MinimalGRPCServer(
            LISTEN_PORT,
            GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS,
            serviceImplRegistrationFunc
        );
        const runResult: Result<null, Error> = await grpcServer.run();
        if (runResult.isErr()) {
            return err(runResult.error);
        }

        return ok(null);
    }

    // ====================================================================================================
    //                                       Private helper functions
    // ====================================================================================================
    private static getEnvVar(envVarName: string, envVarDescription: string): Result<string, Error> {
        if (!(SERIALIZED_CUSTOM_PARAMS_ENV_VAR in process.env)) {
            return err(new Error("Expected an '" + envVarName + "' environment variable containing '" + envVarDescription + "', but none was found"));

        }
        const envVarValue: string = process.env[envVarName]!;
        if (envVarValue === "") {
            return err(new Error("The '" + envVarName + "' environment variable was defined, but is emptystring"));
        }

        return ok(envVarValue);
    }
}

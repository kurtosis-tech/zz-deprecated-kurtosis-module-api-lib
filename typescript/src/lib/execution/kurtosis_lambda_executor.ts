import { ApiContainerServiceClient, NetworkContext } from "kurtosis-core-api-lib";
import { err, ok, Result } from "neverthrow";
import { API_CONTAINER_SOCKET_ENV_VAR, EXECUTION_VOLUME_MOUNTPOINT, SERIALIZED_CUSTOM_PARAMS_ENV_VAR } from "../../kurtosis_lambda_docker_api/kurtosis_lambda_docker_api";
import { KurtosisLambda } from "../kurtosis_lambda/kurtosis_lambda";
import { KurtosisLambdaConfigurator } from "./kurtosis_lambda_configurator";
import * as grpc from 'grpc';
import { ILambdaServiceServer, LambdaServiceService } from "../../kurtosis_lambda_rpc_api_bindings/api_lambda_service_grpc_pb";
import { MinimalGRPCServer, TypedServerOverride } from "minimal-grpc-server";
import { LISTEN_PORT } from "../../kurtosis_lambda_rpc_api_consts/kurtosis_lambda_rpc_api_consts";

const GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS: number = 5;

export class KurtosisLambdaExecutor {
    private readonly configurator: KurtosisLambdaConfigurator;

    constructor(configurator: KurtosisLambdaConfigurator) {
        this.configurator = configurator;
    }

    async run(): Promise<Result<null, Error>> {
        const getSerializedCustomParamsResult: Result<string, Error> = KurtosisLambdaExecutor.getEnvVar(SERIALIZED_CUSTOM_PARAMS_ENV_VAR, "the serialized custom params that the Lambda will consume");
        if (getSerializedCustomParamsResult.isErr()) {
            return err(getSerializedCustomParamsResult.error);
        }
        const serializedCustomParams: string = getSerializedCustomParamsResult.value;

        const createLambdaResult: Result<KurtosisLambda, Error> = this.configurator.parseParamsAndCreateKurtosisLambda(serializedCustomParams);
        if (createLambdaResult.isErr()) {
            return err(createLambdaResult.error);
        }

        const getApiContainerSocketResult: Result<string, Error> = KurtosisLambdaExecutor.getEnvVar(API_CONTAINER_SOCKET_ENV_VAR, "the socket value used in API container connection");
        if (getApiContainerSocketResult.isErr()) {
            return err(getApiContainerSocketResult.error);
        }
        const apiContainerSocket = getApiContainerSocketResult.value;

        // TODO Wrap in exception-handling???
        const apiClient: ApiContainerServiceClient = new ApiContainerServiceClient(apiContainerSocket, grpc.credentials.createInsecure());
        const networkCtx: NetworkContext = new NetworkContext(
            apiClient,
            EXECUTION_VOLUME_MOUNTPOINT
        );

        // TODO
        const lambdaServiceServer: ILambdaServiceServer;
        // lambdaServiceServer := NewKurtosisLambdaServiceServer(lambda, networkCtx)
        const serviceRegistrationFuncs: { (server: TypedServerOverride): void; }[] = [
            (server: TypedServerOverride) => {
                server.addTypedService(LambdaServiceService, lambdaServiceServer);
            }
        ]

        const lambdaServer: MinimalGRPCServer = new MinimalGRPCServer(
            LISTEN_PORT,
            GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS,
            serviceRegistrationFuncs
        );
        const runResult: Result<null, Error> = await lambdaServer.run();
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
        const envVarValue = process.env[envVarName];
        if (envVarValue === "") {
            return err(new Error("The '" + envVarName + "' environment variable was defined, but is emptystring"));
        }

        return ok(envVarValue);
    }
}
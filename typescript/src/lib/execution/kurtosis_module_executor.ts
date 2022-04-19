import { EnclaveContext, getArgsFromEnv } from "kurtosis-core-api-lib";
import { err, ok, Result } from "neverthrow";
import { ExecutableKurtosisModule } from "../kurtosis_modules/executable_kurtosis_modules";
import { KurtosisModuleConfigurator } from "./kurtosis_module_configurator";
import * as grpc from "@grpc/grpc-js";
import { MinimalGRPCServer, TypedServerOverride } from "minimal-grpc-server";
import { ExecutableModuleServiceImpl } from "./executable_module_service_impl";
import { ExecutableModuleServiceService } from "kurtosis-core-api-lib/build/kurtosis_core_rpc_api_bindings/executable_module_service_grpc_pb";

const GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS: number = 5;

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
export class KurtosisModuleExecutor {
    constructor(
        private readonly configurator: KurtosisModuleConfigurator,
    ) {}

    // Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
    public async run(): Promise<Result<null, Error>> {
        const getArgsResult = getArgsFromEnv();
        if (getArgsResult.isErr()) {
            return err(new Error(`An error occurred getting the module container args from the environment`));
        }
        const args = getArgsResult.value;
        const apiContainerSocket = args.apiContainerSocket;
        const serializedCustomParams = args.serializedCustomParams;
        const enclaveId = args.enclaveId;
        const listenPortNum = args.listenPortNum;
        const enclaveDataDirMountpoint = args.enclaveDataDirMountpoint;

        const apiContainerSocketFragments: string[] = apiContainerSocket.split(":");
        const ipAddr: string = apiContainerSocketFragments[0]
        const grpcPortNumStr: string = apiContainerSocketFragments[1]
        const grpcPortNum: number = parseInt(grpcPortNumStr)

        const createModuleResult: Result<ExecutableKurtosisModule, Error> = this.configurator.parseParamsAndCreateExecutableModule(serializedCustomParams);
        if (createModuleResult.isErr()) {
            return err(createModuleResult.error);
        }
        const module: ExecutableKurtosisModule = createModuleResult.value;

        // TODO SECURITY: Use HTTPS to verify we're hitting the correct API container
        const createEnclaveCtxResult = await EnclaveContext.newGrpcNodeEnclaveContext(
            ipAddr,
            grpcPortNum,
            enclaveId,
            enclaveDataDirMountpoint,
        );
        if (createEnclaveCtxResult.isErr()) {
            return err(createEnclaveCtxResult.error);
        }
        const enclaveCtx = createEnclaveCtxResult.value;

        const serviceImpl: ExecutableModuleServiceImpl = new ExecutableModuleServiceImpl(
            module,
            enclaveCtx
        );
        const serviceImplRegistrationFunc: { (server: TypedServerOverride): void; }[] = [
            (server: TypedServerOverride) => {
                server.addTypedService(ExecutableModuleServiceService, serviceImpl);
            }
        ];

        const grpcServer: MinimalGRPCServer = new MinimalGRPCServer(
            listenPortNum,
            GRPC_SERVER_STOP_GRACE_PERIOD_SECONDS,
            serviceImplRegistrationFunc
        );
        const runResult: Result<null, Error> = await grpcServer.runUntilInterrupted();
        if (runResult.isErr()) {
            return err(runResult.error);
        }

        return ok(null);
    }
}

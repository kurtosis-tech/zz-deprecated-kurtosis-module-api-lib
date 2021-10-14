import { ServerUnaryCall, sendUnaryData, ServiceError, Metadata, status, handleCall, Server, ServiceDefinition, GoogleOAuth2Client } from "grpc";
import { NetworkContext } from "kurtosis-core-api-lib";
import { KnownKeysOnly } from "minimal-grpc-server";
import { newExecuteResponse } from "../constructor_calls";
import { ExecutableKurtosisModule } from "../kurtosis_modules/executable_kurtosis_modules";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import { IExecutableModuleServiceServer } from "../../kurtosis_module_rpc_api_bindings/executable_module_service_grpc_pb";
import { ExecuteArgs, ExecuteResponse } from "../../kurtosis_module_rpc_api_bindings/executable_module_service_pb";

// The KnownKeysOnly thing is a workaround due to a silly gRPC requirement that your service implement
//  UnimplementedServiceServer
export class ExecutableModuleServiceImpl implements KnownKeysOnly<IExecutableModuleServiceServer> {
    private readonly module: ExecutableKurtosisModule;
    private readonly networkCtx: NetworkContext;

    constructor(
        module: ExecutableKurtosisModule,
        networkCtx: NetworkContext
    ) {
        this.module = module;
        this.networkCtx = networkCtx;
    }

    public isAvailable(call: ServerUnaryCall<google_protobuf_empty_pb.Empty>, callback: sendUnaryData<google_protobuf_empty_pb.Empty>): void {
        callback(null, new google_protobuf_empty_pb.Empty());
    }

    public execute(call: ServerUnaryCall<ExecuteArgs>, callback: sendUnaryData<ExecuteResponse>): void {
        const args: ExecuteArgs = call.request;
        this.module.execute(
            this.networkCtx, 
            args.getParamsJson()
        ).then(executeResult => {
            if (executeResult.isErr()) {
                callback(executeResult.error, null);
                return;
            }

            const responseJson: string = executeResult.value;
            const executeResponse: ExecuteResponse = newExecuteResponse(responseJson);

            callback(null, executeResponse);
        })
    }
}

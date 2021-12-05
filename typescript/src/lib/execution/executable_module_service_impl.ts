import { ServerUnaryCall, sendUnaryData } from "grpc";
import { EnclaveContext, ExecuteArgs, ExecuteResponse, IExecutableModuleServiceServer } from "kurtosis-core-api-lib";
import { KnownKeysOnly } from "minimal-grpc-server";
import { newExecuteResponse } from "../constructor_calls";
import { ExecutableKurtosisModule } from "../kurtosis_modules/executable_kurtosis_modules";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

// The KnownKeysOnly thing is a workaround due to a silly gRPC requirement that your service implement
//  UnimplementedServiceServer
export class ExecutableModuleServiceImpl implements KnownKeysOnly<IExecutableModuleServiceServer> {
    constructor(
        private readonly module: ExecutableKurtosisModule,
        private readonly enclaveCtx: EnclaveContext,
    ) {}

    public isAvailable(call: ServerUnaryCall<google_protobuf_empty_pb.Empty>, callback: sendUnaryData<google_protobuf_empty_pb.Empty>): void {
        callback(null, new google_protobuf_empty_pb.Empty());
    }

    public execute(call: ServerUnaryCall<ExecuteArgs>, callback: sendUnaryData<ExecuteResponse>): void {
        const args: ExecuteArgs = call.request;
        this.module.execute(
            this.enclaveCtx, 
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

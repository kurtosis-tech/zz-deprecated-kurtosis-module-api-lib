import { ServerUnaryCall, sendUnaryData, ServiceError, Metadata, status, handleCall, Server, ServiceDefinition } from "grpc";
import { NetworkContext } from "kurtosis-core-api-lib";
import { Result } from "neverthrow";
import { ILambdaServiceServer } from "../../kurtosis_lambda_rpc_api_bindings/api_lambda_service_grpc_pb";
import { ExecuteArgs, ExecuteResponse } from "../../kurtosis_lambda_rpc_api_bindings/api_lambda_service_pb";
import { newExecuteResponse } from "../consturctor_calls";
import { KurtosisLambda } from "../kurtosis_lambda/kurtosis_lambda";

class KurtosisLambdaServiceError implements ServiceError {
    readonly code?: status;
    readonly metadata?: Metadata;
    readonly details?: string;
    readonly name: string;
    readonly message: string;
    readonly stack?: string;

    constructor(code: status, from: Error) {
        this.code = code;
        this.metadata = null;
        this.details = null;
        this.name = from.name;
        this.message = from.message;
        this.stack = from.stack;
    }
}

// ====================================== NOTE ========================================================= 
// - https://github.com/agreatfool/grpc_tools_node_protoc_ts/issues/79
// - https://github.com/agreatfool/grpc_tools_node_protoc_ts/blob/master/doc/server_impl_signature.md
// First we need to set up some magic generics inspired by https://github.com/agreatfool/grpc_tools_node_protoc_ts/issues/79#issuecomment-770173789
// ====================================== NOTE ========================================================= 
type KnownKeys<T> = {
    [K in keyof T]: string extends K ? never : number extends K ? never : K
} extends { [_ in keyof T]: infer U } ? U : never;

type KnownOnly<T extends Record<any, any>> = Pick<T, KnownKeys<T>>;

class TypedServerOverride extends Server {
    addTypedService<TypedServiceImplementation extends Record<any,any>>(service: ServiceDefinition, implementation: TypedServiceImplementation): void {
        this.addService(service, implementation);
    }
}

type ITypedLambdaServiceServer = KnownOnly<ILambdaServiceServer>;

export class KurtosisLambdaServiceServer implements ITypedLambdaServiceServer {
    private readonly kurtosisLambda: KurtosisLambda;
    private readonly networkCtx: NetworkContext;

    constructor(
        kurtosisLambda: KurtosisLambda,
        networkCtx: NetworkContext
    ) {
        this.kurtosisLambda = kurtosisLambda;
        this.networkCtx = networkCtx;
    }

    public isAvailable(call: ServerUnaryCall<any>, callback: sendUnaryData<any>): void {
        callback(null, null);
    }

    public execute(call: ServerUnaryCall<ExecuteArgs>, callback: sendUnaryData<ExecuteResponse>): void {
        const args: ExecuteArgs = call.request;
        this.kurtosisLambda.execute(
            this.networkCtx, 
            args.getParamsJson()
        ).then(executeResult => {
            if (executeResult.isErr()) {
                const serviceError: KurtosisLambdaServiceError = new KurtosisLambdaServiceError(
                    status.INTERNAL,
                    executeResult.error
                );
                callback(serviceError, null);
                return
            }

            const responseJson: string = executeResult.value;
            const executeResponse = newExecuteResponse(responseJson)

            callback(null, executeResponse);
        })
    }
}

/*
type KurtosisLambdaServiceServer struct {
	// This embedding is required by gRPC
	kurtosis_lambda_rpc_api_bindings.UnimplementedLambdaServiceServer
	kurtosisLambda kurtosis_lambda.KurtosisLambda
	networkCtx     *networks.NetworkContext
}

func NewKurtosisLambdaServiceServer(kurtosisLambda kurtosis_lambda.KurtosisLambda, networkCtx *networks.NetworkContext) *KurtosisLambdaServiceServer {
	return &KurtosisLambdaServiceServer{kurtosisLambda: kurtosisLambda, networkCtx: networkCtx}
}

func (lambdaService *KurtosisLambdaServiceServer) IsAvailable(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (lambdaService *KurtosisLambdaServiceServer) Execute(ctx context.Context, args *kurtosis_lambda_rpc_api_bindings.ExecuteArgs) (*kurtosis_lambda_rpc_api_bindings.ExecuteResponse, error) {
	result, err := lambdaService.kurtosisLambda.Execute(lambdaService.networkCtx, args.ParamsJson)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred executing the KurtosisLambda")
	}
	executeResponse := &kurtosis_lambda_rpc_api_bindings.ExecuteResponse{
		ResponseJson: result,
	}

	return executeResponse, nil
}
*/

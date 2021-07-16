// GENERATED CODE -- DO NOT EDIT!

// package: api_lambda_service
// file: api_lambda_service.proto

import * as api_lambda_service_pb from "./api_lambda_service_pb";
import * as grpc from "grpc";

interface ILambdaServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  isAvailable: grpc.MethodDefinition<api_lambda_service_pb.IsAvailableArgs, api_lambda_service_pb.IsAvailableResponse>;
  execute: grpc.MethodDefinition<api_lambda_service_pb.ExecuteArgs, api_lambda_service_pb.ExecuteResponse>;
}

export const LambdaServiceService: ILambdaServiceService;

export interface ILambdaServiceServer extends grpc.UntypedServiceImplementation {
  isAvailable: grpc.handleUnaryCall<api_lambda_service_pb.IsAvailableArgs, api_lambda_service_pb.IsAvailableResponse>;
  execute: grpc.handleUnaryCall<api_lambda_service_pb.ExecuteArgs, api_lambda_service_pb.ExecuteResponse>;
}

export class LambdaServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  isAvailable(argument: api_lambda_service_pb.IsAvailableArgs, callback: grpc.requestCallback<api_lambda_service_pb.IsAvailableResponse>): grpc.ClientUnaryCall;
  isAvailable(argument: api_lambda_service_pb.IsAvailableArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.IsAvailableResponse>): grpc.ClientUnaryCall;
  isAvailable(argument: api_lambda_service_pb.IsAvailableArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.IsAvailableResponse>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
}

// GENERATED CODE -- DO NOT EDIT!

// package: kurtosis_lambda_rpc_api_service
// file: api_lambda_service.proto

import * as api_lambda_service_pb from "./api_lambda_service_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "grpc";

interface ILambdaServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  isAvailable: grpc.MethodDefinition<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty>;
  execute: grpc.MethodDefinition<api_lambda_service_pb.ExecuteArgs, api_lambda_service_pb.ExecuteResponse>;
}

export const LambdaServiceService: ILambdaServiceService;

export interface ILambdaServiceServer extends grpc.UntypedServiceImplementation {
  isAvailable: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, google_protobuf_empty_pb.Empty>;
  execute: grpc.handleUnaryCall<api_lambda_service_pb.ExecuteArgs, api_lambda_service_pb.ExecuteResponse>;
}

export class LambdaServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  isAvailable(argument: google_protobuf_empty_pb.Empty, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  isAvailable(argument: google_protobuf_empty_pb.Empty, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  isAvailable(argument: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<google_protobuf_empty_pb.Empty>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
  execute(argument: api_lambda_service_pb.ExecuteArgs, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<api_lambda_service_pb.ExecuteResponse>): grpc.ClientUnaryCall;
}

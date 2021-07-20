// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var api_lambda_service_pb = require('./api_lambda_service_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_lambda_rpc_api_service_ExecuteArgs(arg) {
  if (!(arg instanceof api_lambda_service_pb.ExecuteArgs)) {
    throw new Error('Expected argument of type kurtosis_lambda_rpc_api_service.ExecuteArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_lambda_rpc_api_service_ExecuteArgs(buffer_arg) {
  return api_lambda_service_pb.ExecuteArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_lambda_rpc_api_service_ExecuteResponse(arg) {
  if (!(arg instanceof api_lambda_service_pb.ExecuteResponse)) {
    throw new Error('Expected argument of type kurtosis_lambda_rpc_api_service.ExecuteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_lambda_rpc_api_service_ExecuteResponse(buffer_arg) {
  return api_lambda_service_pb.ExecuteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var LambdaServiceService = exports.LambdaServiceService = {
  // Returns true if the Lambda Module service is available
isAvailable: {
    path: '/kurtosis_lambda_rpc_api_service.LambdaService/IsAvailable',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Executes the main Lambda function in the LambdaService
execute: {
    path: '/kurtosis_lambda_rpc_api_service.LambdaService/Execute',
    requestStream: false,
    responseStream: false,
    requestType: api_lambda_service_pb.ExecuteArgs,
    responseType: api_lambda_service_pb.ExecuteResponse,
    requestSerialize: serialize_kurtosis_lambda_rpc_api_service_ExecuteArgs,
    requestDeserialize: deserialize_kurtosis_lambda_rpc_api_service_ExecuteArgs,
    responseSerialize: serialize_kurtosis_lambda_rpc_api_service_ExecuteResponse,
    responseDeserialize: deserialize_kurtosis_lambda_rpc_api_service_ExecuteResponse,
  },
};

exports.LambdaServiceClient = grpc.makeGenericClientConstructor(LambdaServiceService);

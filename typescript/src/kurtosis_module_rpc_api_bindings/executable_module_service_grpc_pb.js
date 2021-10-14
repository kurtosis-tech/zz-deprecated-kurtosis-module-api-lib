// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var executable_module_service_pb = require('./executable_module_service_pb.js');
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

function serialize_kurtosis_module_rpc_api_service_ExecuteArgs(arg) {
  if (!(arg instanceof executable_module_service_pb.ExecuteArgs)) {
    throw new Error('Expected argument of type kurtosis_module_rpc_api_service.ExecuteArgs');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_module_rpc_api_service_ExecuteArgs(buffer_arg) {
  return executable_module_service_pb.ExecuteArgs.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kurtosis_module_rpc_api_service_ExecuteResponse(arg) {
  if (!(arg instanceof executable_module_service_pb.ExecuteResponse)) {
    throw new Error('Expected argument of type kurtosis_module_rpc_api_service.ExecuteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kurtosis_module_rpc_api_service_ExecuteResponse(buffer_arg) {
  return executable_module_service_pb.ExecuteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// A module that has an "execute" command
var ExecutableModuleServiceService = exports.ExecutableModuleServiceService = {
  // Returns true if the executable module is available
isAvailable: {
    path: '/kurtosis_module_rpc_api_service.ExecutableModuleService/IsAvailable',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
  // Runs the module's execute function
execute: {
    path: '/kurtosis_module_rpc_api_service.ExecutableModuleService/Execute',
    requestStream: false,
    responseStream: false,
    requestType: executable_module_service_pb.ExecuteArgs,
    responseType: executable_module_service_pb.ExecuteResponse,
    requestSerialize: serialize_kurtosis_module_rpc_api_service_ExecuteArgs,
    requestDeserialize: deserialize_kurtosis_module_rpc_api_service_ExecuteArgs,
    responseSerialize: serialize_kurtosis_module_rpc_api_service_ExecuteResponse,
    responseDeserialize: deserialize_kurtosis_module_rpc_api_service_ExecuteResponse,
  },
};

exports.ExecutableModuleServiceClient = grpc.makeGenericClientConstructor(ExecutableModuleServiceService);

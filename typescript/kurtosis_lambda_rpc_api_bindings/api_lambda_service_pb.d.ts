// package: api_lambda_service
// file: api_lambda_service.proto

import * as jspb from "google-protobuf";

export class IsAvailableArgs extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IsAvailableArgs.AsObject;
  static toObject(includeInstance: boolean, msg: IsAvailableArgs): IsAvailableArgs.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IsAvailableArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IsAvailableArgs;
  static deserializeBinaryFromReader(message: IsAvailableArgs, reader: jspb.BinaryReader): IsAvailableArgs;
}

export namespace IsAvailableArgs {
  export type AsObject = {
  }
}

export class IsAvailableResponse extends jspb.Message {
  getIsAvailable(): boolean;
  setIsAvailable(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IsAvailableResponse.AsObject;
  static toObject(includeInstance: boolean, msg: IsAvailableResponse): IsAvailableResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IsAvailableResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IsAvailableResponse;
  static deserializeBinaryFromReader(message: IsAvailableResponse, reader: jspb.BinaryReader): IsAvailableResponse;
}

export namespace IsAvailableResponse {
  export type AsObject = {
    isAvailable: boolean,
  }
}

export class ExecuteArgs extends jspb.Message {
  getParamsJson(): string;
  setParamsJson(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExecuteArgs.AsObject;
  static toObject(includeInstance: boolean, msg: ExecuteArgs): ExecuteArgs.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExecuteArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExecuteArgs;
  static deserializeBinaryFromReader(message: ExecuteArgs, reader: jspb.BinaryReader): ExecuteArgs;
}

export namespace ExecuteArgs {
  export type AsObject = {
    paramsJson: string,
  }
}

export class ExecuteResponse extends jspb.Message {
  getResponseJson(): string;
  setResponseJson(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExecuteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ExecuteResponse): ExecuteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExecuteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExecuteResponse;
  static deserializeBinaryFromReader(message: ExecuteResponse, reader: jspb.BinaryReader): ExecuteResponse;
}

export namespace ExecuteResponse {
  export type AsObject = {
    responseJson: string,
  }
}


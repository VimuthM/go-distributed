// package: runner
// file: runner/runner.proto

import * as jspb from "google-protobuf";

export class RunRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getNodes(): number;
  setNodes(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RunRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RunRequest): RunRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RunRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RunRequest;
  static deserializeBinaryFromReader(message: RunRequest, reader: jspb.BinaryReader): RunRequest;
}

export namespace RunRequest {
  export type AsObject = {
    name: string,
    nodes: number,
  }
}

export class NodeData extends jspb.Message {
  getSenderId(): number;
  setSenderId(value: number): void;

  getReceiverId(): number;
  setReceiverId(value: number): void;

  getRound(): number;
  setRound(value: number): void;

  getMessage(): number;
  setMessage(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NodeData.AsObject;
  static toObject(includeInstance: boolean, msg: NodeData): NodeData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NodeData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NodeData;
  static deserializeBinaryFromReader(message: NodeData, reader: jspb.BinaryReader): NodeData;
}

export namespace NodeData {
  export type AsObject = {
    senderId: number,
    receiverId: number,
    round: number,
    message: number,
  }
}


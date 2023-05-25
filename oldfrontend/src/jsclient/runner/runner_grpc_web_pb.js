/**
 * @fileoverview gRPC-Web generated client stub for runner
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
import * as grpcWeb from 'grpc-web';
grpc.web = grpcWeb;
// grpc.web = require('grpc-web');

const proto = {};
// import * as runner_pb from './runner_pb';
// proto.runner = runner_pb;
// proto.runner = require('./runner_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.runner.RunnerClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.runner.RunnerPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.runner.RunRequest,
 *   !proto.runner.NodeData>}
 */
const methodDescriptor_Runner_runAlgo = new grpc.web.MethodDescriptor(
  '/runner.Runner/runAlgo',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.runner.RunRequest,
  proto.runner.NodeData,
  /**
   * @param {!proto.runner.RunRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.runner.NodeData.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.runner.RunRequest,
 *   !proto.runner.NodeData>}
 */
const methodInfo_Runner_runAlgo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.runner.NodeData,
  /**
   * @param {!proto.runner.RunRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.runner.NodeData.deserializeBinary
);


/**
 * @param {!proto.runner.RunRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.runner.NodeData>}
 *     The XHR Node Readable Stream
 */
proto.runner.RunnerClient.prototype.runAlgo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/runner.Runner/runAlgo',
      request,
      metadata || {},
      methodDescriptor_Runner_runAlgo);
};


/**
 * @param {!proto.runner.RunRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.runner.NodeData>}
 *     The XHR Node Readable Stream
 */
proto.runner.RunnerPromiseClient.prototype.runAlgo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/runner.Runner/runAlgo',
      request,
      metadata || {},
      methodDescriptor_Runner_runAlgo);
};


module.exports = proto.runner;


// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_kardex_pb = require('../v1/kardex_pb.js');

function serialize_v1_RetrieveKardexValuedRequest(arg) {
  if (!(arg instanceof v1_kardex_pb.RetrieveKardexValuedRequest)) {
    throw new Error('Expected argument of type v1.RetrieveKardexValuedRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveKardexValuedRequest(buffer_arg) {
  return v1_kardex_pb.RetrieveKardexValuedRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveKardexValuedResponse(arg) {
  if (!(arg instanceof v1_kardex_pb.RetrieveKardexValuedResponse)) {
    throw new Error('Expected argument of type v1.RetrieveKardexValuedResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveKardexValuedResponse(buffer_arg) {
  return v1_kardex_pb.RetrieveKardexValuedResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var KardexServiceService = exports.KardexServiceService = {
  retrieveKardexValued: {
    path: '/v1.KardexService/RetrieveKardexValued',
    requestStream: false,
    responseStream: false,
    requestType: v1_kardex_pb.RetrieveKardexValuedRequest,
    responseType: v1_kardex_pb.RetrieveKardexValuedResponse,
    requestSerialize: serialize_v1_RetrieveKardexValuedRequest,
    requestDeserialize: deserialize_v1_RetrieveKardexValuedRequest,
    responseSerialize: serialize_v1_RetrieveKardexValuedResponse,
    responseDeserialize: deserialize_v1_RetrieveKardexValuedResponse,
  },
};

exports.KardexServiceClient = grpc.makeGenericClientConstructor(KardexServiceService);

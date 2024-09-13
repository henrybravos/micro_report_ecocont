// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_bank_book_pb = require('../v1/bank_book_pb.js');

function serialize_v1_RetrieveBankBookRequest(arg) {
  if (!(arg instanceof v1_bank_book_pb.RetrieveBankBookRequest)) {
    throw new Error('Expected argument of type v1.RetrieveBankBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveBankBookRequest(buffer_arg) {
  return v1_bank_book_pb.RetrieveBankBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveBankBookResponse(arg) {
  if (!(arg instanceof v1_bank_book_pb.RetrieveBankBookResponse)) {
    throw new Error('Expected argument of type v1.RetrieveBankBookResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveBankBookResponse(buffer_arg) {
  return v1_bank_book_pb.RetrieveBankBookResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var BankBookServiceService = exports.BankBookServiceService = {
  retrieveBankBook: {
    path: '/v1.BankBookService/RetrieveBankBook',
    requestStream: false,
    responseStream: false,
    requestType: v1_bank_book_pb.RetrieveBankBookRequest,
    responseType: v1_bank_book_pb.RetrieveBankBookResponse,
    requestSerialize: serialize_v1_RetrieveBankBookRequest,
    requestDeserialize: deserialize_v1_RetrieveBankBookRequest,
    responseSerialize: serialize_v1_RetrieveBankBookResponse,
    responseDeserialize: deserialize_v1_RetrieveBankBookResponse,
  },
};

exports.BankBookServiceClient = grpc.makeGenericClientConstructor(BankBookServiceService);

// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_cash_book_pb = require('../v1/cash_book_pb.js');

function serialize_v1_RetrieveCashBookRequest(arg) {
  if (!(arg instanceof v1_cash_book_pb.RetrieveCashBookRequest)) {
    throw new Error('Expected argument of type v1.RetrieveCashBookRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveCashBookRequest(buffer_arg) {
  return v1_cash_book_pb.RetrieveCashBookRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveCashBookResponse(arg) {
  if (!(arg instanceof v1_cash_book_pb.RetrieveCashBookResponse)) {
    throw new Error('Expected argument of type v1.RetrieveCashBookResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveCashBookResponse(buffer_arg) {
  return v1_cash_book_pb.RetrieveCashBookResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var CashBookServiceService = exports.CashBookServiceService = {
  retrieveCashBook: {
    path: '/v1.CashBookService/RetrieveCashBook',
    requestStream: false,
    responseStream: false,
    requestType: v1_cash_book_pb.RetrieveCashBookRequest,
    responseType: v1_cash_book_pb.RetrieveCashBookResponse,
    requestSerialize: serialize_v1_RetrieveCashBookRequest,
    requestDeserialize: deserialize_v1_RetrieveCashBookRequest,
    responseSerialize: serialize_v1_RetrieveCashBookResponse,
    responseDeserialize: deserialize_v1_RetrieveCashBookResponse,
  },
};

exports.CashBookServiceClient = grpc.makeGenericClientConstructor(CashBookServiceService);

// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_purchases_pb = require('../v1/purchases_pb.js');

function serialize_v1_RetrievePurchaseReportRequest(arg) {
  if (!(arg instanceof v1_purchases_pb.RetrievePurchaseReportRequest)) {
    throw new Error('Expected argument of type v1.RetrievePurchaseReportRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrievePurchaseReportRequest(buffer_arg) {
  return v1_purchases_pb.RetrievePurchaseReportRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrievePurchaseReportResponse(arg) {
  if (!(arg instanceof v1_purchases_pb.RetrievePurchaseReportResponse)) {
    throw new Error('Expected argument of type v1.RetrievePurchaseReportResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrievePurchaseReportResponse(buffer_arg) {
  return v1_purchases_pb.RetrievePurchaseReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var PurchaseServiceService = exports.PurchaseServiceService = {
  retrievePurchaseReport: {
    path: '/v1.PurchaseService/RetrievePurchaseReport',
    requestStream: false,
    responseStream: false,
    requestType: v1_purchases_pb.RetrievePurchaseReportRequest,
    responseType: v1_purchases_pb.RetrievePurchaseReportResponse,
    requestSerialize: serialize_v1_RetrievePurchaseReportRequest,
    requestDeserialize: deserialize_v1_RetrievePurchaseReportRequest,
    responseSerialize: serialize_v1_RetrievePurchaseReportResponse,
    responseDeserialize: deserialize_v1_RetrievePurchaseReportResponse,
  },
};

exports.PurchaseServiceClient = grpc.makeGenericClientConstructor(PurchaseServiceService);

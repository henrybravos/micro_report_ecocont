// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_sales_pb = require('../v1/sales_pb.js');
var v1_common_pb = require('../v1/common_pb.js');

function serialize_v1_RetrieveSalesPaginatedReportRequest(arg) {
  if (!(arg instanceof v1_sales_pb.RetrieveSalesPaginatedReportRequest)) {
    throw new Error('Expected argument of type v1.RetrieveSalesPaginatedReportRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveSalesPaginatedReportRequest(buffer_arg) {
  return v1_sales_pb.RetrieveSalesPaginatedReportRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveSalesPaginatedReportResponse(arg) {
  if (!(arg instanceof v1_sales_pb.RetrieveSalesPaginatedReportResponse)) {
    throw new Error('Expected argument of type v1.RetrieveSalesPaginatedReportResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveSalesPaginatedReportResponse(buffer_arg) {
  return v1_sales_pb.RetrieveSalesPaginatedReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveSalesResourceReportRequest(arg) {
  if (!(arg instanceof v1_sales_pb.RetrieveSalesResourceReportRequest)) {
    throw new Error('Expected argument of type v1.RetrieveSalesResourceReportRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveSalesResourceReportRequest(buffer_arg) {
  return v1_sales_pb.RetrieveSalesResourceReportRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveSalesResourceReportResponse(arg) {
  if (!(arg instanceof v1_sales_pb.RetrieveSalesResourceReportResponse)) {
    throw new Error('Expected argument of type v1.RetrieveSalesResourceReportResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveSalesResourceReportResponse(buffer_arg) {
  return v1_sales_pb.RetrieveSalesResourceReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var SalesServiceService = exports.SalesServiceService = {
  retrieveSalesPaginatedReport: {
    path: '/v1.SalesService/RetrieveSalesPaginatedReport',
    requestStream: false,
    responseStream: false,
    requestType: v1_sales_pb.RetrieveSalesPaginatedReportRequest,
    responseType: v1_sales_pb.RetrieveSalesPaginatedReportResponse,
    requestSerialize: serialize_v1_RetrieveSalesPaginatedReportRequest,
    requestDeserialize: deserialize_v1_RetrieveSalesPaginatedReportRequest,
    responseSerialize: serialize_v1_RetrieveSalesPaginatedReportResponse,
    responseDeserialize: deserialize_v1_RetrieveSalesPaginatedReportResponse,
  },
  retrieveSalesResourceReport: {
    path: '/v1.SalesService/RetrieveSalesResourceReport',
    requestStream: false,
    responseStream: false,
    requestType: v1_sales_pb.RetrieveSalesResourceReportRequest,
    responseType: v1_sales_pb.RetrieveSalesResourceReportResponse,
    requestSerialize: serialize_v1_RetrieveSalesResourceReportRequest,
    requestDeserialize: deserialize_v1_RetrieveSalesResourceReportRequest,
    responseSerialize: serialize_v1_RetrieveSalesResourceReportResponse,
    responseDeserialize: deserialize_v1_RetrieveSalesResourceReportResponse,
  },
};

exports.SalesServiceClient = grpc.makeGenericClientConstructor(SalesServiceService);

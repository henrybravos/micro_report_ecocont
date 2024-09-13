// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var v1_journal_pb = require('../v1/journal_pb.js');

function serialize_v1_RetrieveGeneralJournalRequest(arg) {
  if (!(arg instanceof v1_journal_pb.RetrieveGeneralJournalRequest)) {
    throw new Error('Expected argument of type v1.RetrieveGeneralJournalRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveGeneralJournalRequest(buffer_arg) {
  return v1_journal_pb.RetrieveGeneralJournalRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveGeneralJournalResponse(arg) {
  if (!(arg instanceof v1_journal_pb.RetrieveGeneralJournalResponse)) {
    throw new Error('Expected argument of type v1.RetrieveGeneralJournalResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveGeneralJournalResponse(buffer_arg) {
  return v1_journal_pb.RetrieveGeneralJournalResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveJournalReportRequest(arg) {
  if (!(arg instanceof v1_journal_pb.RetrieveJournalReportRequest)) {
    throw new Error('Expected argument of type v1.RetrieveJournalReportRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveJournalReportRequest(buffer_arg) {
  return v1_journal_pb.RetrieveJournalReportRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_RetrieveJournalReportResponse(arg) {
  if (!(arg instanceof v1_journal_pb.RetrieveJournalReportResponse)) {
    throw new Error('Expected argument of type v1.RetrieveJournalReportResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_RetrieveJournalReportResponse(buffer_arg) {
  return v1_journal_pb.RetrieveJournalReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var JournalServiceService = exports.JournalServiceService = {
  retrieveJournalReport: {
    path: '/v1.JournalService/RetrieveJournalReport',
    requestStream: false,
    responseStream: false,
    requestType: v1_journal_pb.RetrieveJournalReportRequest,
    responseType: v1_journal_pb.RetrieveJournalReportResponse,
    requestSerialize: serialize_v1_RetrieveJournalReportRequest,
    requestDeserialize: deserialize_v1_RetrieveJournalReportRequest,
    responseSerialize: serialize_v1_RetrieveJournalReportResponse,
    responseDeserialize: deserialize_v1_RetrieveJournalReportResponse,
  },
  retrieveGeneralJournal: {
    path: '/v1.JournalService/RetrieveGeneralJournal',
    requestStream: false,
    responseStream: false,
    requestType: v1_journal_pb.RetrieveGeneralJournalRequest,
    responseType: v1_journal_pb.RetrieveGeneralJournalResponse,
    requestSerialize: serialize_v1_RetrieveGeneralJournalRequest,
    requestDeserialize: deserialize_v1_RetrieveGeneralJournalRequest,
    responseSerialize: serialize_v1_RetrieveGeneralJournalResponse,
    responseDeserialize: deserialize_v1_RetrieveGeneralJournalResponse,
  },
};

exports.JournalServiceClient = grpc.makeGenericClientConstructor(JournalServiceService);

// @generated by protoc-gen-connect-web v0.8.6 with parameter "target=js+dts"
// @generated from file v1/purchases.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { RetrievePurchaseReportRequest, RetrievePurchaseReportResponse } from "./purchases_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service v1.PurchaseService
 */
export declare const PurchaseService: {
  readonly typeName: "v1.PurchaseService",
  readonly methods: {
    /**
     * @generated from rpc v1.PurchaseService.RetrievePurchaseReport
     */
    readonly retrievePurchaseReport: {
      readonly name: "RetrievePurchaseReport",
      readonly I: typeof RetrievePurchaseReportRequest,
      readonly O: typeof RetrievePurchaseReportResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

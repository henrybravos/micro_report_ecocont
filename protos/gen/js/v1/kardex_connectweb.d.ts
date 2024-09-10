// @generated by protoc-gen-connect-web v0.8.6 with parameter "target=js+dts"
// @generated from file v1/kardex.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { RetrieveKardexValuedRequest, RetrieveKardexValuedResponse } from "./kardex_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service v1.KardexService
 */
export declare const KardexService: {
  readonly typeName: "v1.KardexService",
  readonly methods: {
    /**
     * @generated from rpc v1.KardexService.RetrieveKardexValued
     */
    readonly retrieveKardexValued: {
      readonly name: "RetrieveKardexValued",
      readonly I: typeof RetrieveKardexValuedRequest,
      readonly O: typeof RetrieveKardexValuedResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

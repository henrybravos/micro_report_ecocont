// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file v1/cash_book.proto (package v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file v1/cash_book.proto.
 */
export declare const file_v1_cash_book: GenFile;

/**
 * @generated from message v1.AccountBalance
 */
export declare type AccountBalance = Message<"v1.AccountBalance"> & {
  /**
   * @generated from field: string account_id = 1;
   */
  accountId: string;

  /**
   * @generated from field: float amount = 2;
   */
  amount: number;
};

/**
 * Describes the message v1.AccountBalance.
 * Use `create(AccountBalanceSchema)` to create a new message.
 */
export declare const AccountBalanceSchema: GenMessage<AccountBalance>;

/**
 * @generated from message v1.LCash
 */
export declare type LCash = Message<"v1.LCash"> & {
  /**
   * @generated from field: string c_id = 1;
   */
  cId: string;

  /**
   * @generated from field: string c_periodo = 2;
   */
  cPeriodo: string;

  /**
   * @generated from field: string o_cuo_tipo = 3;
   */
  oCuoTipo: string;

  /**
   * @generated from field: string c_identificador = 4;
   */
  cIdentificador: string;

  /**
   * @generated from field: string pc_codigo = 5;
   */
  pcCodigo: string;

  /**
   * @generated from field: string pc_denominacion = 6;
   */
  pcDenominacion: string;

  /**
   * @generated from field: string c_centro_costo_id = 7;
   */
  cCentroCostoId: string;

  /**
   * @generated from field: string tm_codigo = 8;
   */
  tmCodigo: string;

  /**
   * @generated from field: string tc_codigo = 9;
   */
  tcCodigo: string;

  /**
   * @generated from field: string o_serie = 10;
   */
  oSerie: string;

  /**
   * @generated from field: string o_correlativo = 11;
   */
  oCorrelativo: string;

  /**
   * @generated from field: string o_fecha_contable = 12;
   */
  oFechaContable: string;

  /**
   * @generated from field: string o_fecha_vencimiento = 13;
   */
  oFechaVencimiento: string;

  /**
   * @generated from field: string o_fecha_emision = 14;
   */
  oFechaEmision: string;

  /**
   * @generated from field: string pg_fecha = 15;
   */
  pgFecha: string;

  /**
   * @generated from field: string pg_glosa = 16;
   */
  pgGlosa: string;

  /**
   * @generated from field: string o_glosa = 17;
   */
  oGlosa: string;

  /**
   * @generated from field: string o_glosa_referencia = 18;
   */
  oGlosaReferencia: string;

  /**
   * @generated from field: float c_debe = 19;
   */
  cDebe: number;

  /**
   * @generated from field: float c_haber = 20;
   */
  cHaber: number;

  /**
   * @generated from field: string o_dato_estructurado = 21;
   */
  oDatoEstructurado: string;

  /**
   * @generated from field: string c_estado = 22;
   */
  cEstado: string;

  /**
   * @generated from field: string o_estado_le = 23;
   */
  oEstadoLe: string;

  /**
   * @generated from field: string o_observaciones = 24;
   */
  oObservaciones: string;

  /**
   * @generated from field: float o_tipo_cambio = 25;
   */
  oTipoCambio: number;

  /**
   * @generated from field: string o_codigo_libro = 26;
   */
  oCodigoLibro: string;

  /**
   * @generated from field: string o_periodo = 27;
   */
  oPeriodo: string;

  /**
   * @generated from field: float pg_tipo_cambio = 28;
   */
  pgTipoCambio: number;

  /**
   * @generated from field: string pg_fecha_o_fecha_emision = 29;
   */
  pgFechaOFechaEmision: string;
};

/**
 * Describes the message v1.LCash.
 * Use `create(LCashSchema)` to create a new message.
 */
export declare const LCashSchema: GenMessage<LCash>;

/**
 * @generated from message v1.RetrieveCashBookResponse
 */
export declare type RetrieveCashBookResponse = Message<"v1.RetrieveCashBookResponse"> & {
  /**
   * @generated from field: repeated v1.AccountBalance account_balances = 1;
   */
  accountBalances: AccountBalance[];

  /**
   * @generated from field: repeated v1.LCash cash_books = 2;
   */
  cashBooks: LCash[];
};

/**
 * Describes the message v1.RetrieveCashBookResponse.
 * Use `create(RetrieveCashBookResponseSchema)` to create a new message.
 */
export declare const RetrieveCashBookResponseSchema: GenMessage<RetrieveCashBookResponse>;

/**
 * @generated from message v1.RetrieveCashBookRequest
 */
export declare type RetrieveCashBookRequest = Message<"v1.RetrieveCashBookRequest"> & {
  /**
   * @generated from field: string business_id = 1;
   */
  businessId: string;

  /**
   * @generated from field: string period = 2;
   */
  period: string;

  /**
   * @generated from field: repeated string account_ids = 3;
   */
  accountIds: string[];
};

/**
 * Describes the message v1.RetrieveCashBookRequest.
 * Use `create(RetrieveCashBookRequestSchema)` to create a new message.
 */
export declare const RetrieveCashBookRequestSchema: GenMessage<RetrieveCashBookRequest>;

/**
 * @generated from service v1.CashBookService
 */
export declare const CashBookService: GenService<{
  /**
   * @generated from rpc v1.CashBookService.RetrieveCashBook
   */
  retrieveCashBook: {
    methodKind: "unary";
    input: typeof RetrieveCashBookRequestSchema;
    output: typeof RetrieveCashBookResponseSchema;
  },
}>;


// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file v1/bank_book.proto (package v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file v1/bank_book.proto.
 */
export declare const file_v1_bank_book: GenFile;

/**
 * @generated from message v1.BankBalance
 */
export declare type BankBalance = Message<"v1.BankBalance"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: float amount = 2;
   */
  amount: number;
};

/**
 * Describes the message v1.BankBalance.
 * Use `create(BankBalanceSchema)` to create a new message.
 */
export declare const BankBalanceSchema: GenMessage<BankBalance>;

/**
 * @generated from message v1.LBank
 */
export declare type LBank = Message<"v1.LBank"> & {
  /**
   * @generated from field: string b_id = 1;
   */
  bId: string;

  /**
   * @generated from field: string b_periodo = 2;
   */
  bPeriodo: string;

  /**
   * @generated from field: string o_cuo_tipo = 3;
   */
  oCuoTipo: string;

  /**
   * @generated from field: string b_identificador = 4;
   */
  bIdentificador: string;

  /**
   * @generated from field: string tef_codigo = 5;
   */
  tefCodigo: string;

  /**
   * @generated from field: string cf_cuenta = 6;
   */
  cfCuenta: string;

  /**
   * @generated from field: string o_fecha_emion = 7;
   */
  oFechaEmion: string;

  /**
   * @generated from field: string pg_fecha = 8;
   */
  pgFecha: string;

  /**
   * @generated from field: string tmp_codigo = 9;
   */
  tmpCodigo: string;

  /**
   * @generated from field: string pg_glosa = 10;
   */
  pgGlosa: string;

  /**
   * @generated from field: string o_glosa = 11;
   */
  oGlosa: string;

  /**
   * @generated from field: string td_codigo = 12;
   */
  tdCodigo: string;

  /**
   * @generated from field: string p_numero = 13;
   */
  pNumero: string;

  /**
   * @generated from field: string p_razon_social = 14;
   */
  pRazonSocial: string;

  /**
   * @generated from field: string b_numero_transaccion = 15;
   */
  bNumeroTransaccion: string;

  /**
   * @generated from field: float b_debe = 16;
   */
  bDebe: number;

  /**
   * @generated from field: float b_haber = 17;
   */
  bHaber: number;

  /**
   * @generated from field: string b_estado = 18;
   */
  bEstado: string;

  /**
   * @generated from field: string pc_codigo = 19;
   */
  pcCodigo: string;

  /**
   * @generated from field: string pc_denominacion = 20;
   */
  pcDenominacion: string;

  /**
   * @generated from field: string o_estado_le = 21;
   */
  oEstadoLe: string;

  /**
   * @generated from field: string o_observaciones = 22;
   */
  oObservaciones: string;

  /**
   * @generated from field: float o_tipo_cambio = 23;
   */
  oTipoCambio: number;

  /**
   * @generated from field: float pg_tipo_cambio = 24;
   */
  pgTipoCambio: number;

  /**
   * @generated from field: string pg_fecha_o_fecha_emision = 25;
   */
  pgFechaOFechaEmision: string;
};

/**
 * Describes the message v1.LBank.
 * Use `create(LBankSchema)` to create a new message.
 */
export declare const LBankSchema: GenMessage<LBank>;

/**
 * @generated from message v1.RetrieveBankBookResponse
 */
export declare type RetrieveBankBookResponse = Message<"v1.RetrieveBankBookResponse"> & {
  /**
   * @generated from field: repeated v1.BankBalance bank_balances = 1;
   */
  bankBalances: BankBalance[];

  /**
   * @generated from field: repeated v1.LBank bank_books = 2;
   */
  bankBooks: LBank[];
};

/**
 * Describes the message v1.RetrieveBankBookResponse.
 * Use `create(RetrieveBankBookResponseSchema)` to create a new message.
 */
export declare const RetrieveBankBookResponseSchema: GenMessage<RetrieveBankBookResponse>;

/**
 * @generated from message v1.RetrieveBankBookRequest
 */
export declare type RetrieveBankBookRequest = Message<"v1.RetrieveBankBookRequest"> & {
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
 * Describes the message v1.RetrieveBankBookRequest.
 * Use `create(RetrieveBankBookRequestSchema)` to create a new message.
 */
export declare const RetrieveBankBookRequestSchema: GenMessage<RetrieveBankBookRequest>;

/**
 * @generated from service v1.BankBookService
 */
export declare const BankBookService: GenService<{
  /**
   * @generated from rpc v1.BankBookService.RetrieveBankBook
   */
  retrieveBankBook: {
    methodKind: "unary";
    input: typeof RetrieveBankBookRequestSchema;
    output: typeof RetrieveBankBookResponseSchema;
  },
}>;

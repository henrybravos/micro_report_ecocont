// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file v1/cash_book.proto (package v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";

/**
 * Describes the file v1/cash_book.proto.
 */
export const file_v1_cash_book = /*@__PURE__*/
  fileDesc("ChJ2MS9jYXNoX2Jvb2sucHJvdG8SAnYxIjQKDkFjY291bnRCYWxhbmNlEhIKCmFjY291bnRfaWQYASABKAkSDgoGYW1vdW50GAIgASgCIoUFCgVMQ2FzaBIMCgRjX2lkGAEgASgJEhEKCWNfcGVyaW9kbxgCIAEoCRISCgpvX2N1b190aXBvGAMgASgJEhcKD2NfaWRlbnRpZmljYWRvchgEIAEoCRIRCglwY19jb2RpZ28YBSABKAkSFwoPcGNfZGVub21pbmFjaW9uGAYgASgJEhkKEWNfY2VudHJvX2Nvc3RvX2lkGAcgASgJEhEKCXRtX2NvZGlnbxgIIAEoCRIRCgl0Y19jb2RpZ28YCSABKAkSDwoHb19zZXJpZRgKIAEoCRIVCg1vX2NvcnJlbGF0aXZvGAsgASgJEhgKEG9fZmVjaGFfY29udGFibGUYDCABKAkSGwoTb19mZWNoYV92ZW5jaW1pZW50bxgNIAEoCRIXCg9vX2ZlY2hhX2VtaXNpb24YDiABKAkSEAoIcGdfZmVjaGEYDyABKAkSEAoIcGdfZ2xvc2EYECABKAkSDwoHb19nbG9zYRgRIAEoCRIaChJvX2dsb3NhX3JlZmVyZW5jaWEYEiABKAkSDgoGY19kZWJlGBMgASgCEg8KB2NfaGFiZXIYFCABKAISGwoTb19kYXRvX2VzdHJ1Y3R1cmFkbxgVIAEoCRIQCghjX2VzdGFkbxgWIAEoCRITCgtvX2VzdGFkb19sZRgXIAEoCRIXCg9vX29ic2VydmFjaW9uZXMYGCABKAkSFQoNb190aXBvX2NhbWJpbxgZIAEoAhIWCg5vX2NvZGlnb19saWJybxgaIAEoCRIRCglvX3BlcmlvZG8YGyABKAkSFgoOcGdfdGlwb19jYW1iaW8YHCABKAISIAoYcGdfZmVjaGFfb19mZWNoYV9lbWlzaW9uGB0gASgJImcKGFJldHJpZXZlQ2FzaEJvb2tSZXNwb25zZRIsChBhY2NvdW50X2JhbGFuY2VzGAEgAygLMhIudjEuQWNjb3VudEJhbGFuY2USHQoKY2FzaF9ib29rcxgCIAMoCzIJLnYxLkxDYXNoIlMKF1JldHJpZXZlQ2FzaEJvb2tSZXF1ZXN0EhMKC2J1c2luZXNzX2lkGAEgASgJEg4KBnBlcmlvZBgCIAEoCRITCgthY2NvdW50X2lkcxgDIAMoCTJgCg9DYXNoQm9va1NlcnZpY2USTQoQUmV0cmlldmVDYXNoQm9vaxIbLnYxLlJldHJpZXZlQ2FzaEJvb2tSZXF1ZXN0GhwudjEuUmV0cmlldmVDYXNoQm9va1Jlc3BvbnNlQnQKBmNvbS52MUINQ2FzaEJvb2tQcm90b1ABWjNnaXRodWIuY29tL2hlbnJ5YnJhdm8vbWljcm8tcmVwb3J0L3Byb3Rvcy9nZW4vZ28vdjGiAgNWWFiqAgJWMcoCAlYx4gIOVjFcR1BCTWV0YWRhdGHqAgJWMWIGcHJvdG8z");

/**
 * Describes the message v1.AccountBalance.
 * Use `create(AccountBalanceSchema)` to create a new message.
 */
export const AccountBalanceSchema = /*@__PURE__*/
  messageDesc(file_v1_cash_book, 0);

/**
 * Describes the message v1.LCash.
 * Use `create(LCashSchema)` to create a new message.
 */
export const LCashSchema = /*@__PURE__*/
  messageDesc(file_v1_cash_book, 1);

/**
 * Describes the message v1.RetrieveCashBookResponse.
 * Use `create(RetrieveCashBookResponseSchema)` to create a new message.
 */
export const RetrieveCashBookResponseSchema = /*@__PURE__*/
  messageDesc(file_v1_cash_book, 2);

/**
 * Describes the message v1.RetrieveCashBookRequest.
 * Use `create(RetrieveCashBookRequestSchema)` to create a new message.
 */
export const RetrieveCashBookRequestSchema = /*@__PURE__*/
  messageDesc(file_v1_cash_book, 3);

/**
 * @generated from service v1.CashBookService
 */
export const CashBookService = /*@__PURE__*/
  serviceDesc(file_v1_cash_book, 0);

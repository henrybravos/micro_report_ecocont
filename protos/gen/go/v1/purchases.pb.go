// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/purchases.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PurchaseReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OId                     string `protobuf:"bytes,1,opt,name=o_id,json=oId,proto3" json:"o_id,omitempty"`
	OPeriodo                string `protobuf:"bytes,2,opt,name=o_periodo,json=oPeriodo,proto3" json:"o_periodo,omitempty"`
	OCuo                    string `protobuf:"bytes,3,opt,name=o_cuo,json=oCuo,proto3" json:"o_cuo,omitempty"`
	CIdentificadorLinea     string `protobuf:"bytes,4,opt,name=c_identificador_linea,json=cIdentificadorLinea,proto3" json:"c_identificador_linea,omitempty"`
	OFechaEmision           string `protobuf:"bytes,5,opt,name=o_fecha_emision,json=oFechaEmision,proto3" json:"o_fecha_emision,omitempty"`
	OFechaVencimiento       string `protobuf:"bytes,6,opt,name=o_fecha_vencimiento,json=oFechaVencimiento,proto3" json:"o_fecha_vencimiento,omitempty"`
	TcoCodigo               string `protobuf:"bytes,7,opt,name=tco_codigo,json=tcoCodigo,proto3" json:"tco_codigo,omitempty"`
	OSerie                  string `protobuf:"bytes,8,opt,name=o_serie,json=oSerie,proto3" json:"o_serie,omitempty"`
	CDua                    string `protobuf:"bytes,9,opt,name=c_dua,json=cDua,proto3" json:"c_dua,omitempty"`
	OCorrelativo            string `protobuf:"bytes,10,opt,name=o_correlativo,json=oCorrelativo,proto3" json:"o_correlativo,omitempty"`
	CNumeroFinal            string `protobuf:"bytes,11,opt,name=c_numero_final,json=cNumeroFinal,proto3" json:"c_numero_final,omitempty"`
	TdCodigo                string `protobuf:"bytes,12,opt,name=td_codigo,json=tdCodigo,proto3" json:"td_codigo,omitempty"`
	PNumero                 string `protobuf:"bytes,13,opt,name=p_numero,json=pNumero,proto3" json:"p_numero,omitempty"`
	PRazonSocial            string `protobuf:"bytes,14,opt,name=p_razon_social,json=pRazonSocial,proto3" json:"p_razon_social,omitempty"`
	CBase_1                 string `protobuf:"bytes,15,opt,name=c_base_1,json=cBase1,proto3" json:"c_base_1,omitempty"`
	CIgv_1                  string `protobuf:"bytes,16,opt,name=c_igv_1,json=cIgv1,proto3" json:"c_igv_1,omitempty"`
	CBase_2                 string `protobuf:"bytes,17,opt,name=c_base_2,json=cBase2,proto3" json:"c_base_2,omitempty"`
	CIgv_2                  string `protobuf:"bytes,18,opt,name=c_igv_2,json=cIgv2,proto3" json:"c_igv_2,omitempty"`
	CBase_3                 string `protobuf:"bytes,19,opt,name=c_base_3,json=cBase3,proto3" json:"c_base_3,omitempty"`
	CIgv_3                  string `protobuf:"bytes,20,opt,name=c_igv_3,json=cIgv3,proto3" json:"c_igv_3,omitempty"`
	CNoGravada              string `protobuf:"bytes,21,opt,name=c_no_gravada,json=cNoGravada,proto3" json:"c_no_gravada,omitempty"`
	CIsc                    string `protobuf:"bytes,22,opt,name=c_isc,json=cIsc,proto3" json:"c_isc,omitempty"`
	COtros                  string `protobuf:"bytes,23,opt,name=c_otros,json=cOtros,proto3" json:"c_otros,omitempty"`
	CTotal                  string `protobuf:"bytes,24,opt,name=c_total,json=cTotal,proto3" json:"c_total,omitempty"`
	TmCodigo                string `protobuf:"bytes,25,opt,name=tm_codigo,json=tmCodigo,proto3" json:"tm_codigo,omitempty"`
	OTipoCambio             string `protobuf:"bytes,26,opt,name=o_tipo_cambio,json=oTipoCambio,proto3" json:"o_tipo_cambio,omitempty"`
	CFechaCdpm              string `protobuf:"bytes,27,opt,name=c_fecha_cdpm,json=cFechaCdpm,proto3" json:"c_fecha_cdpm,omitempty"`
	TicoCodigo              string `protobuf:"bytes,28,opt,name=tico_codigo,json=ticoCodigo,proto3" json:"tico_codigo,omitempty"`
	CSerieCdpm              string `protobuf:"bytes,29,opt,name=c_serie_cdpm,json=cSerieCdpm,proto3" json:"c_serie_cdpm,omitempty"`
	CDuaCdpm                string `protobuf:"bytes,30,opt,name=c_dua_cdpm,json=cDuaCdpm,proto3" json:"c_dua_cdpm,omitempty"`
	CCorrelativoCdpm        string `protobuf:"bytes,31,opt,name=c_correlativo_cdpm,json=cCorrelativoCdpm,proto3" json:"c_correlativo_cdpm,omitempty"`
	CFechaDetraccion        string `protobuf:"bytes,32,opt,name=c_fecha_detraccion,json=cFechaDetraccion,proto3" json:"c_fecha_detraccion,omitempty"`
	CNumeroDetraccion       string `protobuf:"bytes,33,opt,name=c_numero_detraccion,json=cNumeroDetraccion,proto3" json:"c_numero_detraccion,omitempty"`
	CMarcaCdpRetencion      string `protobuf:"bytes,34,opt,name=c_marca_cdp_retencion,json=cMarcaCdpRetencion,proto3" json:"c_marca_cdp_retencion,omitempty"`
	BbssCodigo              string `protobuf:"bytes,35,opt,name=bbss_codigo,json=bbssCodigo,proto3" json:"bbss_codigo,omitempty"`
	CContrato               string `protobuf:"bytes,36,opt,name=c_contrato,json=cContrato,proto3" json:"c_contrato,omitempty"`
	CError_1                string `protobuf:"bytes,37,opt,name=c_error_1,json=cError1,proto3" json:"c_error_1,omitempty"`
	CError_2                string `protobuf:"bytes,38,opt,name=c_error_2,json=cError2,proto3" json:"c_error_2,omitempty"`
	CError_3                string `protobuf:"bytes,39,opt,name=c_error_3,json=cError3,proto3" json:"c_error_3,omitempty"`
	CError_4                string `protobuf:"bytes,40,opt,name=c_error_4,json=cError4,proto3" json:"c_error_4,omitempty"`
	CMedioPago              string `protobuf:"bytes,41,opt,name=c_medio_pago,json=cMedioPago,proto3" json:"c_medio_pago,omitempty"`
	OImporteCdpRegimenSunat string `protobuf:"bytes,42,opt,name=o_importe_cdp_regimen_sunat,json=oImporteCdpRegimenSunat,proto3" json:"o_importe_cdp_regimen_sunat,omitempty"`
	OTipoCdpRegimenSunat    string `protobuf:"bytes,43,opt,name=o_tipo_cdp_regimen_sunat,json=oTipoCdpRegimenSunat,proto3" json:"o_tipo_cdp_regimen_sunat,omitempty"`
	CIcbper                 string `protobuf:"bytes,44,opt,name=c_icbper,json=cIcbper,proto3" json:"c_icbper,omitempty"`
	CEstado                 string `protobuf:"bytes,45,opt,name=c_estado,json=cEstado,proto3" json:"c_estado,omitempty"`
}

func (x *PurchaseReport) Reset() {
	*x = PurchaseReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_purchases_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseReport) ProtoMessage() {}

func (x *PurchaseReport) ProtoReflect() protoreflect.Message {
	mi := &file_v1_purchases_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseReport.ProtoReflect.Descriptor instead.
func (*PurchaseReport) Descriptor() ([]byte, []int) {
	return file_v1_purchases_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseReport) GetOId() string {
	if x != nil {
		return x.OId
	}
	return ""
}

func (x *PurchaseReport) GetOPeriodo() string {
	if x != nil {
		return x.OPeriodo
	}
	return ""
}

func (x *PurchaseReport) GetOCuo() string {
	if x != nil {
		return x.OCuo
	}
	return ""
}

func (x *PurchaseReport) GetCIdentificadorLinea() string {
	if x != nil {
		return x.CIdentificadorLinea
	}
	return ""
}

func (x *PurchaseReport) GetOFechaEmision() string {
	if x != nil {
		return x.OFechaEmision
	}
	return ""
}

func (x *PurchaseReport) GetOFechaVencimiento() string {
	if x != nil {
		return x.OFechaVencimiento
	}
	return ""
}

func (x *PurchaseReport) GetTcoCodigo() string {
	if x != nil {
		return x.TcoCodigo
	}
	return ""
}

func (x *PurchaseReport) GetOSerie() string {
	if x != nil {
		return x.OSerie
	}
	return ""
}

func (x *PurchaseReport) GetCDua() string {
	if x != nil {
		return x.CDua
	}
	return ""
}

func (x *PurchaseReport) GetOCorrelativo() string {
	if x != nil {
		return x.OCorrelativo
	}
	return ""
}

func (x *PurchaseReport) GetCNumeroFinal() string {
	if x != nil {
		return x.CNumeroFinal
	}
	return ""
}

func (x *PurchaseReport) GetTdCodigo() string {
	if x != nil {
		return x.TdCodigo
	}
	return ""
}

func (x *PurchaseReport) GetPNumero() string {
	if x != nil {
		return x.PNumero
	}
	return ""
}

func (x *PurchaseReport) GetPRazonSocial() string {
	if x != nil {
		return x.PRazonSocial
	}
	return ""
}

func (x *PurchaseReport) GetCBase_1() string {
	if x != nil {
		return x.CBase_1
	}
	return ""
}

func (x *PurchaseReport) GetCIgv_1() string {
	if x != nil {
		return x.CIgv_1
	}
	return ""
}

func (x *PurchaseReport) GetCBase_2() string {
	if x != nil {
		return x.CBase_2
	}
	return ""
}

func (x *PurchaseReport) GetCIgv_2() string {
	if x != nil {
		return x.CIgv_2
	}
	return ""
}

func (x *PurchaseReport) GetCBase_3() string {
	if x != nil {
		return x.CBase_3
	}
	return ""
}

func (x *PurchaseReport) GetCIgv_3() string {
	if x != nil {
		return x.CIgv_3
	}
	return ""
}

func (x *PurchaseReport) GetCNoGravada() string {
	if x != nil {
		return x.CNoGravada
	}
	return ""
}

func (x *PurchaseReport) GetCIsc() string {
	if x != nil {
		return x.CIsc
	}
	return ""
}

func (x *PurchaseReport) GetCOtros() string {
	if x != nil {
		return x.COtros
	}
	return ""
}

func (x *PurchaseReport) GetCTotal() string {
	if x != nil {
		return x.CTotal
	}
	return ""
}

func (x *PurchaseReport) GetTmCodigo() string {
	if x != nil {
		return x.TmCodigo
	}
	return ""
}

func (x *PurchaseReport) GetOTipoCambio() string {
	if x != nil {
		return x.OTipoCambio
	}
	return ""
}

func (x *PurchaseReport) GetCFechaCdpm() string {
	if x != nil {
		return x.CFechaCdpm
	}
	return ""
}

func (x *PurchaseReport) GetTicoCodigo() string {
	if x != nil {
		return x.TicoCodigo
	}
	return ""
}

func (x *PurchaseReport) GetCSerieCdpm() string {
	if x != nil {
		return x.CSerieCdpm
	}
	return ""
}

func (x *PurchaseReport) GetCDuaCdpm() string {
	if x != nil {
		return x.CDuaCdpm
	}
	return ""
}

func (x *PurchaseReport) GetCCorrelativoCdpm() string {
	if x != nil {
		return x.CCorrelativoCdpm
	}
	return ""
}

func (x *PurchaseReport) GetCFechaDetraccion() string {
	if x != nil {
		return x.CFechaDetraccion
	}
	return ""
}

func (x *PurchaseReport) GetCNumeroDetraccion() string {
	if x != nil {
		return x.CNumeroDetraccion
	}
	return ""
}

func (x *PurchaseReport) GetCMarcaCdpRetencion() string {
	if x != nil {
		return x.CMarcaCdpRetencion
	}
	return ""
}

func (x *PurchaseReport) GetBbssCodigo() string {
	if x != nil {
		return x.BbssCodigo
	}
	return ""
}

func (x *PurchaseReport) GetCContrato() string {
	if x != nil {
		return x.CContrato
	}
	return ""
}

func (x *PurchaseReport) GetCError_1() string {
	if x != nil {
		return x.CError_1
	}
	return ""
}

func (x *PurchaseReport) GetCError_2() string {
	if x != nil {
		return x.CError_2
	}
	return ""
}

func (x *PurchaseReport) GetCError_3() string {
	if x != nil {
		return x.CError_3
	}
	return ""
}

func (x *PurchaseReport) GetCError_4() string {
	if x != nil {
		return x.CError_4
	}
	return ""
}

func (x *PurchaseReport) GetCMedioPago() string {
	if x != nil {
		return x.CMedioPago
	}
	return ""
}

func (x *PurchaseReport) GetOImporteCdpRegimenSunat() string {
	if x != nil {
		return x.OImporteCdpRegimenSunat
	}
	return ""
}

func (x *PurchaseReport) GetOTipoCdpRegimenSunat() string {
	if x != nil {
		return x.OTipoCdpRegimenSunat
	}
	return ""
}

func (x *PurchaseReport) GetCIcbper() string {
	if x != nil {
		return x.CIcbper
	}
	return ""
}

func (x *PurchaseReport) GetCEstado() string {
	if x != nil {
		return x.CEstado
	}
	return ""
}

type RetrievePurchaseReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Period     string `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *RetrievePurchaseReportRequest) Reset() {
	*x = RetrievePurchaseReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_purchases_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievePurchaseReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievePurchaseReportRequest) ProtoMessage() {}

func (x *RetrievePurchaseReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_purchases_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievePurchaseReportRequest.ProtoReflect.Descriptor instead.
func (*RetrievePurchaseReportRequest) Descriptor() ([]byte, []int) {
	return file_v1_purchases_proto_rawDescGZIP(), []int{1}
}

func (x *RetrievePurchaseReportRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *RetrievePurchaseReportRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

type RetrievePurchaseReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PurchaseReport `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RetrievePurchaseReportResponse) Reset() {
	*x = RetrievePurchaseReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_purchases_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievePurchaseReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievePurchaseReportResponse) ProtoMessage() {}

func (x *RetrievePurchaseReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_purchases_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievePurchaseReportResponse.ProtoReflect.Descriptor instead.
func (*RetrievePurchaseReportResponse) Descriptor() ([]byte, []int) {
	return file_v1_purchases_proto_rawDescGZIP(), []int{2}
}

func (x *RetrievePurchaseReportResponse) GetData() []*PurchaseReport {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_purchases_proto protoreflect.FileDescriptor

var file_v1_purchases_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xd7, 0x0b, 0x0a, 0x0e, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x6f, 0x12, 0x13, 0x0a, 0x05, 0x6f,
	0x5f, 0x63, 0x75, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x43, 0x75, 0x6f,
	0x12, 0x32, 0x0a, 0x15, 0x63, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x64, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x6f, 0x72, 0x4c,
	0x69, 0x6e, 0x65, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f,
	0x65, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x46, 0x65, 0x63, 0x68, 0x61, 0x45, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x6f, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x76, 0x65, 0x6e, 0x63, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x46, 0x65, 0x63, 0x68,
	0x61, 0x56, 0x65, 0x6e, 0x63, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x63, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x63, 0x6f, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x5f, 0x64, 0x75, 0x61, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x44, 0x75, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x5f, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6f, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x6f, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67,
	0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x64, 0x43, 0x6f, 0x64, 0x69, 0x67,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x5f, 0x72, 0x61, 0x7a, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x52, 0x61, 0x7a, 0x6f, 0x6e, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x12, 0x18, 0x0a, 0x08, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x31, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x42, 0x61, 0x73, 0x65, 0x31, 0x12, 0x16, 0x0a, 0x07,
	0x63, 0x5f, 0x69, 0x67, 0x76, 0x5f, 0x31, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x49, 0x67, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x08, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x32,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x42, 0x61, 0x73, 0x65, 0x32, 0x12, 0x16,
	0x0a, 0x07, 0x63, 0x5f, 0x69, 0x67, 0x76, 0x5f, 0x32, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x49, 0x67, 0x76, 0x32, 0x12, 0x18, 0x0a, 0x08, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x33, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x42, 0x61, 0x73, 0x65, 0x33,
	0x12, 0x16, 0x0a, 0x07, 0x63, 0x5f, 0x69, 0x67, 0x76, 0x5f, 0x33, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x49, 0x67, 0x76, 0x33, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x5f, 0x6e, 0x6f,
	0x5f, 0x67, 0x72, 0x61, 0x76, 0x61, 0x64, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x4e, 0x6f, 0x47, 0x72, 0x61, 0x76, 0x61, 0x64, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x5f,
	0x69, 0x73, 0x63, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x49, 0x73, 0x63, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x5f, 0x6f, 0x74, 0x72, 0x6f, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x4f, 0x74, 0x72, 0x6f, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6d, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x22,
	0x0a, 0x0d, 0x6f, 0x5f, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x62,
	0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x63, 0x64,
	0x70, 0x6d, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x46, 0x65, 0x63, 0x68, 0x61,
	0x43, 0x64, 0x70, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6f, 0x5f, 0x63, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6f, 0x43,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x5f, 0x63, 0x64, 0x70, 0x6d, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x43, 0x64, 0x70, 0x6d, 0x12, 0x1c, 0x0a, 0x0a, 0x63, 0x5f, 0x64, 0x75, 0x61,
	0x5f, 0x63, 0x64, 0x70, 0x6d, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x44, 0x75,
	0x61, 0x43, 0x64, 0x70, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x6f, 0x5f, 0x63, 0x64, 0x70, 0x6d, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x6f, 0x43,
	0x64, 0x70, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x5f, 0x66, 0x65, 0x63, 0x68, 0x61, 0x5f, 0x64,
	0x65, 0x74, 0x72, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x46, 0x65, 0x63, 0x68, 0x61, 0x44, 0x65, 0x74, 0x72, 0x61, 0x63, 0x63, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x5f, 0x64, 0x65,
	0x74, 0x72, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x63, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x44, 0x65, 0x74, 0x72, 0x61, 0x63, 0x63, 0x69, 0x6f,
	0x6e, 0x12, 0x31, 0x0a, 0x15, 0x63, 0x5f, 0x6d, 0x61, 0x72, 0x63, 0x61, 0x5f, 0x63, 0x64, 0x70,
	0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x63, 0x4d, 0x61, 0x72, 0x63, 0x61, 0x43, 0x64, 0x70, 0x52, 0x65, 0x74, 0x65, 0x6e,
	0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x62, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x69, 0x67, 0x6f, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x62, 0x73, 0x73, 0x43,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x09, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x31, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x31,
	0x12, 0x1a, 0x0a, 0x09, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x32, 0x18, 0x26, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x12, 0x1a, 0x0a, 0x09,
	0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x33, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x33, 0x12, 0x1a, 0x0a, 0x09, 0x63, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x34, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x34, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x6f, 0x5f,
	0x70, 0x61, 0x67, 0x6f, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x4d, 0x65, 0x64,
	0x69, 0x6f, 0x50, 0x61, 0x67, 0x6f, 0x12, 0x3c, 0x0a, 0x1b, 0x6f, 0x5f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x5f, 0x63, 0x64, 0x70, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6d, 0x65, 0x6e, 0x5f,
	0x73, 0x75, 0x6e, 0x61, 0x74, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6f, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x43, 0x64, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6d, 0x65, 0x6e, 0x53,
	0x75, 0x6e, 0x61, 0x74, 0x12, 0x36, 0x0a, 0x18, 0x6f, 0x5f, 0x74, 0x69, 0x70, 0x6f, 0x5f, 0x63,
	0x64, 0x70, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6d, 0x65, 0x6e, 0x5f, 0x73, 0x75, 0x6e, 0x61, 0x74,
	0x18, 0x2b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x64, 0x70,
	0x52, 0x65, 0x67, 0x69, 0x6d, 0x65, 0x6e, 0x53, 0x75, 0x6e, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x5f, 0x69, 0x63, 0x62, 0x70, 0x65, 0x72, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x49, 0x63, 0x62, 0x70, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x5f, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x22, 0x58, 0x0a, 0x1d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x48, 0x0a, 0x1e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x72, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x6e, 0x72, 0x79, 0x62, 0x72, 0x61, 0x76, 0x6f, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58,
	0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_purchases_proto_rawDescOnce sync.Once
	file_v1_purchases_proto_rawDescData = file_v1_purchases_proto_rawDesc
)

func file_v1_purchases_proto_rawDescGZIP() []byte {
	file_v1_purchases_proto_rawDescOnce.Do(func() {
		file_v1_purchases_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_purchases_proto_rawDescData)
	})
	return file_v1_purchases_proto_rawDescData
}

var file_v1_purchases_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_purchases_proto_goTypes = []any{
	(*PurchaseReport)(nil),                 // 0: v1.PurchaseReport
	(*RetrievePurchaseReportRequest)(nil),  // 1: v1.RetrievePurchaseReportRequest
	(*RetrievePurchaseReportResponse)(nil), // 2: v1.RetrievePurchaseReportResponse
}
var file_v1_purchases_proto_depIdxs = []int32{
	0, // 0: v1.RetrievePurchaseReportResponse.data:type_name -> v1.PurchaseReport
	1, // 1: v1.PurchaseService.RetrievePurchaseReport:input_type -> v1.RetrievePurchaseReportRequest
	2, // 2: v1.PurchaseService.RetrievePurchaseReport:output_type -> v1.RetrievePurchaseReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_purchases_proto_init() }
func file_v1_purchases_proto_init() {
	if File_v1_purchases_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_purchases_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_purchases_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RetrievePurchaseReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_purchases_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RetrievePurchaseReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_purchases_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_purchases_proto_goTypes,
		DependencyIndexes: file_v1_purchases_proto_depIdxs,
		MessageInfos:      file_v1_purchases_proto_msgTypes,
	}.Build()
	File_v1_purchases_proto = out.File
	file_v1_purchases_proto_rawDesc = nil
	file_v1_purchases_proto_goTypes = nil
	file_v1_purchases_proto_depIdxs = nil
}
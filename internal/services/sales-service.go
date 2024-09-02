package services

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/henrybravo/micro-report/internal/report/excel"
	"github.com/henrybravo/micro-report/internal/report/pdf"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/pkg/validate"
	v1 "github.com/henrybravo/micro-report/protos/gen/go/v1"
	"log"
)

type SalesServer struct {
	SalesRepo    *repo.SalesRepository
	BusinessRepo *repo.BusinessRepository
}

func (s *SalesServer) RetrieveSalesPaginatedReport(
	_ context.Context,
	req *connect.Request[v1.RetrieveSalesPaginatedReportRequest],
) (*connect.Response[v1.RetrieveSalesPaginatedReportResponse], error) {
	log.Println("Request headers: ", req.Header())
	businessID := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	pageSize := req.Msg.GetPageSize()
	page := req.Msg.GetPage()
	paginationDefault := GetPaginationOrDefault(page, pageSize)
	if !validate.IsValidUUID(businessID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}

	sales, pagination, err := s.SalesRepo.GetSalesReports(businessID, period, paginationDefault)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	pagination.Page = paginationDefault.Page
	pagination.PageSize = paginationDefault.PageSize
	pagination.Offset = paginationDefault.Offset
	res := connect.NewResponse(&v1.RetrieveSalesPaginatedReportResponse{
		Sales:      sales,
		Pagination: pagination,
	})
	res.Header().Set("Report-Version", "v1")
	return res, nil
}

func (s *SalesServer) RetrieveSalesResourceReport(
	_ context.Context,
	req *connect.Request[v1.RetrieveSalesResourceReportRequest],
) (*connect.Response[v1.RetrieveSalesResourceReportResponse], error) {
	log.Println("Request headers: ", req.Header())
	businessID := req.Msg.GetBusinessId()
	period := req.Msg.GetPeriod()
	typeResource := req.Msg.GetType()
	if !validate.IsValidUUID(businessID) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid business ID"))
	}
	if !validate.IsValidPeriod(period) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid period"))
	}
	if typeResource == v1.TypeResource_TYPE_RESOURCE_UNSPECIFIED {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid type resource"))
	}

	sales, _, err := s.SalesRepo.GetSalesReports(businessID, period, nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	business, err := s.BusinessRepo.GetBusinessByID(businessID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	var pathName string
	if typeResource == v1.TypeResource_TYPE_RESOURCE_XLSX {
		generatorXLSX := excel.NewSalesGenerator()
		pathName, err = generatorXLSX.GenerateSalesReport(business, period, sales)
	}
	if typeResource == v1.TypeResource_TYPE_RESOURCE_PDF {
		generatorPdf := pdf.NewSalesGenerator()
		pathName, err = generatorPdf.GenerateSalesReport(business, sales, period)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&v1.RetrieveSalesResourceReportResponse{
		Url: pathName,
	})
	res.Header().Set("Report-Version", "v1")
	return res, err
}

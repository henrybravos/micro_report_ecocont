package services

import (
	"bytes"
	"github.com/henrybravo/micro-report/internal/report"
	"github.com/henrybravo/micro-report/internal/repositories"
)

type ReportService struct {
	SalesRepo      *repositories.SalesRepository
	ExcelGenerator *report.ExcelGenerator
}

func NewReportService(salesRepo *repositories.SalesRepository, excelGenerator *report.ExcelGenerator) *ReportService {
	return &ReportService{
		SalesRepo:      salesRepo,
		ExcelGenerator: excelGenerator,
	}
}

func (s *ReportService) CreateExcelSales(companyID, period string) (*bytes.Buffer, error) {
	sales, err := s.SalesRepo.GetSalesReports(companyID, period)
	if err != nil {
		return nil, err
	}
	excel, err := s.ExcelGenerator.GenerateSalesReport(sales)
	if err != nil {
		return nil, err
	}
	return excel, nil
}

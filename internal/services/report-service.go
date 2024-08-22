package services

import (
	"bytes"
	"github.com/henrybravo/micro-report/internal/report"
	repo "github.com/henrybravo/micro-report/internal/repositories"
)

type ReportService struct {
	SalesRepo      *repo.SalesRepository
	ExcelGenerator *report.ExcelGenerator
}

func NewReportService(salesRepo *repo.SalesRepository, excelGenerator *report.ExcelGenerator) *ReportService {
	return &ReportService{
		SalesRepo:      salesRepo,
		ExcelGenerator: excelGenerator,
	}
}

func (s *ReportService) CreateExcelSales(companyID, period string) (*bytes.Buffer, error) {
	sales, _, err := s.SalesRepo.GetSalesReports(companyID, period, repo.PaginationParams{Pagination: false})
	if err != nil {
		return nil, err
	}
	excel, err := s.ExcelGenerator.GenerateSalesReport(sales)
	if err != nil {
		return nil, err
	}
	return excel, nil
}
func (s *ReportService) CreateSalesPaginated(companyID, period string, offset, pageSize int) ([]repo.SalesReport, *repo.Pagination, error) {
	sales, pagination, err := s.SalesRepo.GetSalesReports(companyID, period, repo.PaginationParams{Pagination: true, Offset: offset, Limit: pageSize})
	if err != nil {
		return nil, nil, err
	}
	return sales, pagination, nil
}

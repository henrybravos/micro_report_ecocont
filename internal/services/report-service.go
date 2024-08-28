package services

import (
	"bytes"
	"fmt"
	"github.com/henrybravo/micro-report/internal/report"
	repo "github.com/henrybravo/micro-report/internal/repositories"
	"time"
)

type ReportService struct {
	salesRepo      *repo.SalesRepository
	businessRepo   *repo.BusinessRepository
	excelGenerator *report.ExcelGenerator
	pdfGenerator   *report.PdfGenerator
}

func NewReportService(salesRepo *repo.SalesRepository, businessRepo *repo.BusinessRepository, excelGenerator *report.ExcelGenerator, pdfGenerator *report.PdfGenerator) *ReportService {
	return &ReportService{
		salesRepo:      salesRepo,
		businessRepo:   businessRepo,
		excelGenerator: excelGenerator,
		pdfGenerator:   pdfGenerator,
	}
}
func (s *ReportService) GetBusinessByID(id string) (*repo.Business, error) {
	business, err := s.businessRepo.GetBusinessByID(id)
	if err != nil {
		return nil, err
	}
	return business, nil
}

func (s *ReportService) CreateSalesPaginated(businessID, period string, isPaginated bool, offset, pageSize int) ([]repo.SalesReport, *repo.Pagination, error) {
	sales, pagination, err := s.salesRepo.GetSalesReports(businessID, period, repo.PaginationParams{Pagination: isPaginated, Offset: offset, Limit: pageSize})
	if err != nil {
		return nil, nil, err
	}
	return sales, pagination, nil
}
func (s *ReportService) CreateExcelSales(businessID, period string) (*bytes.Buffer, error) {
	start := time.Now()
	sales, _, err := s.salesRepo.GetSalesReports(businessID, period, repo.PaginationParams{Pagination: false})
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)
	fmt.Printf("Consulta a la base de datos completada en %v\n", duration)
	fmt.Println("Procediendo a generar el EXCEL")

	excel, err := s.excelGenerator.GenerateSalesReport(sales)
	if err != nil {
		return nil, err
	}
	return excel, nil
}
func (s *ReportService) CreatePDFSales(businessID, period string) (*bytes.Buffer, error) {
	start := time.Now()
	business, err := s.businessRepo.GetBusinessByID(businessID)
	sales, _, err := s.salesRepo.GetSalesReports(businessID, period, repo.PaginationParams{Pagination: false})
	duration := time.Since(start)
	fmt.Printf("Consulta a la base de datos completada en %v, total: %d items\n", duration, len(sales))
	fmt.Println("Procediendo a generar el PDF")
	start = time.Now()
	if err != nil {
		return nil, err
	}
	pdf, err := s.pdfGenerator.GeneratePDF(*business, sales, period)
	duration = time.Since(start)
	fmt.Printf("Generación del PDF completada en %v\n", duration)
	if err != nil {
		return nil, err
	}
	return pdf, nil
}

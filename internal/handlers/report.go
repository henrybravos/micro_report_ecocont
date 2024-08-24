package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henrybravo/micro-report/internal/report"
	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/internal/services"
	"net/http"
	"strconv"
	"time"
)

type ReportHandler struct {
	ReportService *services.ReportService
}

func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{ReportService: reportService}
}

func (h *ReportHandler) GetSalesExcelReport(c *gin.Context) {
	companyID := c.Query("companyID")
	period := c.Query("period")

	excelBuffer, err := h.ReportService.CreateExcelSales(companyID, period)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Set headers to indicate an Excel file download
	fileName := generateFileName("sales-report", "xlsx")

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Length", fmt.Sprintf("%d", excelBuffer.Len()))

	// Write the buffer to the response body
	index, err := c.Writer.Write(excelBuffer.Bytes())
	if err != nil {
		fmt.Println("Error writing response index", index)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

}
func generateFileName(baseName, ext string) string {
	currentDate := time.Now()
	return fmt.Sprintf("%s-%s.%s", baseName, currentDate.Format("20060102150405"), ext)
}
func (h *ReportHandler) GetSalesReportPaginated(c *gin.Context) {
	companyID := c.Query("companyID")
	period := c.Query("period")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	// Default values if not provided
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 15
	}
	// Calculate the offset for the database query
	offset := (page - 1) * pageSize

	sales, pagination, err := h.ReportService.CreateSalesPaginated(companyID, period, true, offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Page = page
	pagination.PageSize = pageSize
	if sales == nil {
		sales = []repositories.SalesReport{}
	}
	c.JSON(http.StatusOK, gin.H{
		"pagination": pagination,
		"data":       sales,
	})

}

func (h *ReportHandler) GetSalesReportPDF(c *gin.Context) {
	companyID := c.Query("companyID")
	period := c.Query("period")
	sales, _, err := h.ReportService.CreateSalesPaginated(companyID, period, false, 0, 0)
	buffer, err := report.GeneratePDF(sales)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fileName := generateFileName("sales-report", "pdf")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", fileName))
	c.Header("Content-Length", fmt.Sprintf("%d", buffer.Len()))

	index, err := c.Writer.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Error writing response index", index)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

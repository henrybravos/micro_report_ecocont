package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	fileName := generateFileName()

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
func generateFileName() string {
	currentDate := time.Now()
	return fmt.Sprintf("Report-sales-%s.xlsx", currentDate.Format("20060102150405"))
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

	sales, pagination, err := h.ReportService.CreateSalesPaginated(companyID, period, offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pagination.Page = page
	pagination.PageSize = pageSize
	c.JSON(http.StatusOK, gin.H{
		"pagination": pagination,
		"sales":      sales,
	})

}

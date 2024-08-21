package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henrybravo/micro-report/internal/services"
	"time"
)

type ReportHandler struct {
	ReportService *services.ReportService
}

func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{ReportService: reportService}
}

func (h *ReportHandler) GetSalesReport(c *gin.Context) {
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

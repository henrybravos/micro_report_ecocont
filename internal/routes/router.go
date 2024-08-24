package routes

import (
	"github.com/gin-gonic/gin"
	hand "github.com/henrybravo/micro-report/internal/handlers"
)

func SetupRouter(reportHandler *hand.ReportHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/api/sales-excel", reportHandler.GetSalesExcelReport)
	router.GET("/api/sales-pdf", reportHandler.GetSalesReportPDF)
	router.GET("/api/sales-paginated", reportHandler.GetSalesReportPaginated)
	return router
}

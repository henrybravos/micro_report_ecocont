package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henrybravo/micro-report/internal/handlers"
)

func SetupRouter(reportHandler *handlers.ReportHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/api/sales-excel", reportHandler.GetSalesReport)
	return router
}

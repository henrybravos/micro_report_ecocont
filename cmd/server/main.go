package main

import (
	"github.com/henrybravo/micro-report/internal/handlers"
	"github.com/henrybravo/micro-report/internal/report"
	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/internal/routes"
	"github.com/henrybravo/micro-report/internal/services"
	"github.com/henrybravo/micro-report/pkg/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pathDB := db.GetDatabaseURL()
	dbConnection, err := db.ConnectToDB(pathDB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConnection.Pool.Close()
	salesRepo := repositories.NewSalesRepository(dbConnection)
	businessRepo := repositories.NewBusinessRepository(dbConnection)
	excelGenerator := report.NewExcelGenerator()
	pdfGenerator := report.NewPdfGenerator()
	reportService := services.NewReportService(salesRepo, businessRepo, excelGenerator, pdfGenerator)
	reportHandler := handlers.NewReportHandler(reportService)
	router := routes.SetupRouter(reportHandler)
	log.Println("Server listening on port 8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

}

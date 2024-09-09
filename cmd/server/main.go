package main

import (
	"github.com/henrybravo/micro-report/internal/repositories"
	"github.com/henrybravo/micro-report/internal/services"
	"github.com/henrybravo/micro-report/pkg/db"
	"github.com/henrybravo/micro-report/protos/gen/go/v1/v1connect"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	mux := http.NewServeMux()

	salesRepo := repositories.NewSalesRepository(dbConnection)
	businessRepo := repositories.NewBusinessRepository(dbConnection)
	journalRepo := repositories.NewJournalRepository(dbConnection)
	cashRepo := repositories.NewCashRepository(dbConnection)
	bankBookRepo := repositories.NewBankBookRepository(dbConnection)

	salesServer := &services.SalesServer{
		SalesRepo:    salesRepo,
		BusinessRepo: businessRepo,
	}
	path, handler := v1connect.NewSalesServiceHandler(salesServer)
	mux.Handle(path, handler)
	journalServer := &services.JournalServer{
		JournalRepo: journalRepo,
	}
	path, handler = v1connect.NewJournalServiceHandler(journalServer)
	mux.Handle(path, handler)

	cashServer := &services.CashServer{
		CashRepo: cashRepo,
	}
	path, handler = v1connect.NewCashBookServiceHandler(cashServer)
	mux.Handle(path, handler)

	bankBookServer := &services.BookServer{
		BookRepo: bankBookRepo,
	}
	path, handler = v1connect.NewBankBookServiceHandler(bankBookServer)
	mux.Handle(path, handler)

	mux.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("tmp"))))
	err = http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

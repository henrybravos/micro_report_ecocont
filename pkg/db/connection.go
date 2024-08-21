package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

type Connection struct {
	Pool *pgxpool.Pool
}

// ConnectToDB establece una conexión a la base de datos PostgreSQL usando pgx
func ConnectToDB(databaseURL string) (*Connection, error) {
	// Configurar la conexión a la base de datos
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}

	// Crear el pool de conexiones
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	// Probar la conexión
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &Connection{Pool: pool}, nil
}

func GetDatabaseURL() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	log.Println("Connecting to database at", databaseURL)
	return databaseURL
}

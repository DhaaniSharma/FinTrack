package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	databaseURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}
	DB = pool
	fmt.Println("Connected to PostgreSQL database successfully.")

	return nil
}

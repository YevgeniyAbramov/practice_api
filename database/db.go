package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	DB = pool
	fmt.Println("Connected to database")

}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

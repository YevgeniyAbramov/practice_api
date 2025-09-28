package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	DB = pool
	fmt.Println("Connected to database")

}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

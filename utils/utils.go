package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env файл не найден")
	}
}

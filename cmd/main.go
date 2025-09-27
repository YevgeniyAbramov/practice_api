package main

import (
	"log"
	"practice_api/database"
	"practice_api/routes"
	"practice_api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.Init()

	database.ConnectDB()
	defer database.CloseDB()

	app := fiber.New()
	routes.Routes(app)

	log.Fatal(app.Listen(":8010"))
}

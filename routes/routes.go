package routes

import (
	"practice_api/handlers"

	_ "practice_api/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Routes(app *fiber.App) {
	// Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	v1 := app.Group("/v1/")

	// Ручка статуса
	app.Get("/status", handlers.GetStatus)

	// API v1
	user := v1.Group("/users")
	{
		user.Get("/", handlers.GetUsers)
		user.Post("/", handlers.CreateUser)
		user.Put("/update/:id", handlers.UpdateUser)
		user.Delete("/soft-delete/:id", handlers.SoftDeleteUser)
	}
}

package handlers

import "github.com/gofiber/fiber/v2"

// CreateUser godoc
// @Router /v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "ok",
	})
}

// GetStatus godoc
// @Router /status [get]
func GetStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "active",
	})
}

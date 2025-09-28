package handlers

import (
	"practice_api/database"
	"practice_api/models"
	"practice_api/response"

	"github.com/gofiber/fiber/v2"
)

// GetStatus godoc
// @Tags check status
// @Description Проверяет статус API
// @Router /status [get]
func GetStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "active",
	})
}

// CreateUser godoc
// @Summary Создать пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "Данные пользователя"
// @Success 200 {object} response.APIResponse{result=models.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	resp, err := database.CreateUser(req)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
		Result:  resp,
	})
}

// GetUsers godoc
// @Tags users
// @Summary Получить пользователей
// @Description Получает список всех пользователей
// @Produce json
// @Success 200 {object} response.APIResponse{result=[]models.User}
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/users [get]
func GetUsers(c *fiber.Ctx) error {

	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
		Result:  users,
	})
}

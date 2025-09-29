package handlers

import (
	"practice_api/database"
	"practice_api/models"
	"practice_api/response"
	"strconv"

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

// UpdateUser godoc
// @Summary Обновить пользователя
// @Description Обновляет пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body models.UpdateUserRequest true "Данные пользователя"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/users/update/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: "id is required",
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: "invalid id",
		})
	}

	var req models.UpdateUserRequest
	err = c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = req.Validate()
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = database.UpdateUser(req, userId)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
	})
}

// SoftDeleteUser godoc
// @Summary Удалить пользователя
// @Description Помечает пользователя как удаленого
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/users/soft-delete/{id} [delete]
func SoftDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: "id is required",
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: "invalid id",
		})
	}

	err = database.SoftDeleteUser(userId)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
	})

}

// RestoreUser godoc
// @Summary Восстановить пользователя
// @Description Восстанавливает пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/users/restore/{id} [patch]
func RestoreUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: "user id is required",
		})
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err = database.RestoreUser(userId)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
	})

}

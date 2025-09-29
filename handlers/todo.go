package handlers

import (
	"practice_api/database"
	"practice_api/models"
	"practice_api/response"

	"github.com/gofiber/fiber/v2"
)

// CreateToDo godoc
// @Summary Создать задачу
// @Description Создает новую задачу
// @Tags ToDo
// @Accept json
// @Produce json
// @Param user body models.CreateToDoReq true "Данные пользователя"
// @Success 200 {object} response.APIResponse{result=models.ToDo}
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/todo/create [post]
func CreateToDo(c *fiber.Ctx) error {
	var req models.CreateToDoReq

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	result, err := database.CreateToDo(&req)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Message: "ok",
		Result:  result,
	})

}

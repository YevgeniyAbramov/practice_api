package handlers

import (
	"practice_api/database"
	"practice_api/models"
	"practice_api/response"
	"strconv"

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

// GetToDo godoc
// @Summary Получить список задач
// @Description Получает список задач
// @Tags ToDo
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} response.APIResponse{result=[]models.ToDo}
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /v1/todo/get/{id} [get]
func GetToDo(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	result, err := database.GetToDo(userID)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Status:  true,
		Result:  result,
		Message: "ok",
	})

}

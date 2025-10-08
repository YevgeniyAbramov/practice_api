package database

import (
	"fmt"
	"practice_api/models"
	"strings"
)

func CreateToDo(todo *models.CreateToDoReq) (*models.ToDo, error) {
	query := `
		INSERT INTO auth.todos (user_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, name, description, created_at, updated_at`

	var newToDo models.ToDo
	err := DB.QueryRow(
		query,
		todo.UserId, todo.Name, todo.Description,
	).Scan(
		&newToDo.Id,
		&newToDo.UserId,
		&newToDo.Name,
		&newToDo.Description,
		&newToDo.CreatedAt,
		&newToDo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &newToDo, nil
}

func GetToDo(userID int) ([]models.ToDo, error) {
	query := `SELECT * FROM auth.todos WHERE user_id = $1 AND deleted_at IS NULL`

	var todos []models.ToDo
	err := DB.Select(&todos, query, userID)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func UpdateToDo(id int, todo *models.UpdateToDoReq) (*models.ToDo, error) {
	query := `UPDATE auth.todos SET `
	args := []interface{}{}
	i := 1

	if todo.Name != "" {
		query += fmt.Sprintf("name = $%d", i)
		args = append(args, todo.Name)
		i++
	}

	if todo.Description != nil {
		query += fmt.Sprintf("description = $%d", i)
		args = append(args, todo.Description)
		i++
	}

	if todo.Status != nil {
		query += fmt.Sprintf("status = $%d", i)
		args = append(args, todo.Status)
		i++
	}

	query = strings.TrimSuffix(query, ", ")
	query += fmt.Sprintf(" WHERE id = $%d RETURNING *", i)
	args = append(args, id)

	var update models.ToDo

	err := DB.QueryRowx(query, args...).StructScan(&update)
	if err != nil {
		return nil, err
	}

	return &update, err

}

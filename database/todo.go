package database

import "practice_api/models"

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

package database

import (
	"fmt"
	"practice_api/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	err := DB.Select(&users, "SELECT * FROM auth.users WHERE deleted_at is NULL")
	if err != nil {
		return users, err
	}

	return users, nil
}

func CreateUser(user models.CreateUserRequest) (models.User, error) {
	query := `
		INSERT INTO auth.users (login, first_name,last_name)
		VALUES ($1,$2,$3)
		RETURNING *`

	var newUser models.User

	err := DB.Get(&newUser, query, user.Login, user.FirstName, user.LastName)
	if err != nil {
		return models.User{}, err
	}

	return newUser, nil

}

func UpdateUser(user models.UpdateUserRequest, userId int) error {
	query := `
		UPDATE auth.users
		SET login = $1, first_name = $2, last_name = $3, updated_at = NOW()
		where id = $4
		AND deleted_at IS NULL`

	res, err := DB.Exec(query, user.Login, user.FirstName, user.LastName, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d or user is soft-deleted", userId)
	}

	return nil
}

func SoftDeleteUser(id int) error {
	query := `
		UPDATE auth.users
		SET deleted_at = NOW()
		WHERE id = $1`

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

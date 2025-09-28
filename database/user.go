package database

import (
	"practice_api/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	err := DB.Select(&users, "SELECT * FROM auth.users")
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

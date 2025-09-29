package models

import (
	"fmt"
	"time"
)

// User пользователь системы
//
// swagger:model User
type User struct {
	ID        int        `db:"id" json:"id"`                           // ID пользователя
	Login     string     `db:"login" json:"login"`                     // Логин пользователя
	FirstName *string    `db:"first_name" json:"first_name,omitempty"` // Имя пользователя
	LastName  *string    `db:"last_name" json:"last_name,omitempty"`   // Фамилия пользователя
	CreatedAt time.Time  `db:"created_at" json:"created_at"`           // Дата создания
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`           // Дата обновления
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"` // Дата удаления (если есть)
}

// CreateUserRequest тело запроса на создание пользователя
// swagger:model CreateUserRequest
type CreateUserRequest struct {
	Login     string  `json:"login" example:"johndoe"`
	FirstName *string `json:"first_name,omitempty" example:"John"`
	LastName  *string `json:"last_name,omitempty" example:"Doe"`
}

// CreateUserRequest тело запроса на создание пользователя
type UpdateUserRequest struct {
	Login     string `json:"login" example:"johndoe"`
	FirstName string `json:"first_name,omitempty" example:"John"`
	LastName  string `json:"last_name,omitempty" example:"Doe"`
}

func (u *UpdateUserRequest) Validate() error {
	if u.Login == "" {
		return fmt.Errorf("login is required")
	}

	if u.FirstName == "" {
		return fmt.Errorf("first name is required")
	}

	if u.LastName == "" {
		return fmt.Errorf("last name is required")
	}

	return nil
}

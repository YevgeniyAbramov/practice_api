package models

import "time"

type TodoStatus int

const (
	StatusNew TodoStatus = iota
	StatusInProgress
	StatusDone
)

// swagger:model ToDo
type ToDo struct {
	Id          int        `json:"id" db:"id"`
	UserId      int        `json:"user_id" db:"user_id"`                   // ID пользователя задачи
	Name        string     `json:"name" db:"name"`                         // Название задачи
	Description *string    `json:"description,omitempty" db:"description"` //Описание задачи
	Status      TodoStatus `json:"status" db:"status"`                     //Статус
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`             // Дата создания
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`             // Дата обновления
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`   // Дата удаления (если есть)
}

// swagger:model CreateToDoReq
type CreateToDoReq struct {
	UserId      int     `json:"user_id" db:"user_id"`                   // ID пользователя задачи
	Name        string  `json:"name" db:"name"`                         // Название задачи
	Description *string `json:"description,omitempty" db:"description"` //Описание задачи
}

package models

import "time"

type User struct {
	ID        int        `db:"id" json:"id"`
	Login     string     `db:"login" json:"login"`
	FirstName *string    `db:"first_name" json:"first_name,omitempty"`
	LastName  *string    `db:"last_name" json:"last_name,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

package models

import "time"

type Todo struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Completed bool      `json:"completed" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"title"`
	UpdatedAt time.Time `json:"updated_at" db:"title"`
	UserID    string    `json:"user_id" db:"title"`
}

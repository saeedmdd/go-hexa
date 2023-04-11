package models

import "time"

type TaskModel struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Color       string    `json:"color,omitempty"`
	StartsAt    time.Time `json:"starts_at"`
	DoneAt      time.Time `json:"done_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

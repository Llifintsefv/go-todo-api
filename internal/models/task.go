package models

import "time"

type Task struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

func IsValidStatus(status string) bool {
	switch status {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}

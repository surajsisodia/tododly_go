package models

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"task_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"last_updated_at"`
	UpdatedBy   string    `json:"last_updated_by"`
}

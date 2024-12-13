package models

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey" json:"user_id"`
	FirstName    *string   `json:"first_name,omitempty"`
	LastName     *string   `json:"last_name,omitempty"`
	PhotoUrl     *string   `json:"photo_url,omitempty"`
	Email        string    `gorm:"not null" json:"email"`
	Token        *string   `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	CreatedBy    string    `gorm:"not null" json:"created_by"`
	UpdatedAt    time.Time `gorm:"not null" json:"last_updated_at"`
	UpdatedBy    string    `gorm:"not null" json:"last_updated_by"`
}

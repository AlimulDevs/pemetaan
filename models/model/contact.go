package model

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID          string         `json:"id" gorm:"primaryKey;notNull;size:255"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

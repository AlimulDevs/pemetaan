package model

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID          string         `json:"id" gorm:"primaryKey;notNull;size:255"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

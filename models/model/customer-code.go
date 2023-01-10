package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCode struct {
	ID        string         `json:"id" gorm:"primaryKey;notNull;size:255"`
	Email     string         `json:"email"`
	Code      string         `json:"code" gorm:"notNull;size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

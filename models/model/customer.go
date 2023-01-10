package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           string         `json:"id" gorm:"primaryKey;notNull;size:255"`
	Name         string         `json:"name" gorm:"notNull;size:255"`
	Email        string         `json:"email" gorm:"notNull;unique;size:255"`
	Password     string         `json:"password" gorm:"notNull"`
	ProfileImage string         `json:"profile_image" gorm:"size:255;default:null"`
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

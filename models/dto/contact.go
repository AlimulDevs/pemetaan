package dto

type Contact struct {
	ID          string `json:"id" gorm:"primaryKey;notNull;size:255"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
}

type ContactTransaction struct {
	ID          string `json:"id" gorm:"primaryKey;notNull;size:255"`
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

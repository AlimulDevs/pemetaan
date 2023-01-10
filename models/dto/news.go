package dto

type News struct {
	ID          string `json:"id" gorm:"primaryKey;notNull;size:255"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
type NewsTransaction struct {
	ID          string `json:"id" gorm:"primaryKey;notNull;size:255"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

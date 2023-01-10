package newsrepository

import (
	"golang/models/dto"
)

type Repository interface {
	GetAll() ([]dto.News, error)
	GetByID(id string) (dto.News, error)
	Create(input dto.NewsTransaction) error
	Update(id string, input dto.NewsTransaction) error
	Delete(id string) error
}

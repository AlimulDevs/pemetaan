package contactrepository

import (
	"golang/models/dto"
)

type Repository interface {
	GetAll() ([]dto.Contact, error)
	GetByID(id string) (dto.Contact, error)
	Create(input dto.ContactTransaction) error
	Update(id string, input dto.ContactTransaction) error
	Delete(id string) error
}

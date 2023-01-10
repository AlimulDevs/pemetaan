package contactservice

import (
	"golang/helper"
	"golang/models/dto"
	contactrepository "golang/repository/contactRepository"
)

type Service interface {
	GetAll() ([]dto.Contact, error)
	GetByID(id string) (dto.Contact, error)
	Create(input dto.ContactTransaction) error
	Update(id string, input dto.ContactTransaction) error
	Delete(id string) error
}

type service struct {
	repo contactrepository.Repository
}

func (sv *service) GetAll() ([]dto.Contact, error) {
	response, err := sv.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (sv *service) GetByID(id string) (dto.Contact, error) {
	response, err := sv.repo.GetByID(id)
	if err != nil {
		return dto.Contact{}, err
	}
	return response, nil
}

func (sv *service) Create(input dto.ContactTransaction) error {
	id := helper.GenerateUUID()
	input.ID = id
	err := sv.repo.Create(input)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) Update(id string, input dto.ContactTransaction) error {
	err := sv.repo.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) Delete(id string) error {
	err := sv.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(repo contactrepository.Repository) Service {
	return &service{
		repo: repo,
	}
}

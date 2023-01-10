package newsservice

import (
	"golang/helper"
	"golang/models/dto"
	newsrepository "golang/repository/newsRepository"
)

type Service interface {
	GetAll() ([]dto.News, error)
	GetByID(id string) (dto.News, error)
	Create(input dto.NewsTransaction) error
	Update(id string, input dto.NewsTransaction) error
	Delete(id string) error
}

type service struct {
	repo newsrepository.Repository
}

func (sv *service) GetAll() ([]dto.News, error) {
	response, err := sv.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (sv *service) GetByID(id string) (dto.News, error) {
	response, err := sv.repo.GetByID(id)
	if err != nil {
		return dto.News{}, err
	}
	return response, nil
}

func (sv *service) Create(input dto.NewsTransaction) error {
	id := helper.GenerateUUID()
	input.ID = id
	err := sv.repo.Create(input)
	if err != nil {
		return err
	}
	return nil
}

func (sv *service) Update(id string, input dto.NewsTransaction) error {
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

func NewService(repo newsrepository.Repository) Service {
	return &service{
		repo: repo,
	}
}

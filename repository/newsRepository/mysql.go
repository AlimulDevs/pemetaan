package newsrepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (config *repository) GetAll() ([]dto.News, error) {
	var dataModel []model.News
	var response []dto.News
	err := config.db.Find(&dataModel).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&response, dataModel)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (config *repository) GetByID(id string) (dto.News, error) {
	var dataModel model.News
	var response dto.News

	err := config.db.Where("id=?", id).First(&dataModel).Error
	if err != nil {
		return dto.News{}, err
	}
	err = copier.Copy(&response, dataModel)
	if err != nil {
		return dto.News{}, err
	}

	return response, nil
}

func (config *repository) Create(input dto.NewsTransaction) error {
	var dataModel model.News
	err := copier.Copy(&dataModel, &input)
	if err != nil {
		return err
	}
	err = config.db.Create(&dataModel).Error
	if err != nil {
		return err
	}

	return nil

}

func (config *repository) Update(id string, input dto.NewsTransaction) error {
	var dataModel model.News
	err := copier.Copy(&dataModel, &input)
	if err != nil {
		return err
	}
	err = config.db.Where("id=?", id).Updates(&dataModel).Error
	if err != nil {
		return err
	}

	return nil

}

func (config *repository) Delete(id string) error {

	err := config.db.Where("id=?", id).Unscoped().Delete(&model.News{})
	if err != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

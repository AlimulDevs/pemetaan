package contactrepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (config *repository) GetAll() ([]dto.Contact, error) {
	var dataModel []model.Contact
	var response []dto.Contact
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

func (config *repository) GetByID(id string) (dto.Contact, error) {
	var dataModel model.Contact
	var response dto.Contact

	err := config.db.Where("id=?", id).First(&dataModel).Error
	if err != nil {
		return dto.Contact{}, err
	}
	err = copier.Copy(&response, dataModel)
	if err != nil {
		return dto.Contact{}, err
	}

	return response, nil
}

func (config *repository) Create(input dto.ContactTransaction) error {
	var dataModel model.Contact
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

func (config *repository) Update(id string, input dto.ContactTransaction) error {
	var dataModel model.Contact
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

	err := config.db.Where("id=?", id).Unscoped().Delete(&model.Contact{})
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

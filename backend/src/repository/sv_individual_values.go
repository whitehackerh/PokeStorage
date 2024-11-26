package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVIndividualValuesRepository interface {
		Create(*gorm.DB, model.SVIndividualValues) error
		Update(*gorm.DB, model.SVIndividualValues) error
		Delete(*gorm.DB, string) error
	}
	SVIndividualValuesRepository struct {
		Db *gorm.DB
	}
)

func NewSVIndividualValuesRepository(db *gorm.DB) ISVIndividualValuesRepository {
	return &SVIndividualValuesRepository{
		Db: db,
	}
}

func (s *SVIndividualValuesRepository) Create(tx *gorm.DB, model model.SVIndividualValues) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVIndividualValuesRepository) Update(tx *gorm.DB, model model.SVIndividualValues) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVIndividualValuesRepository) Delete(tx *gorm.DB, id string) error {
	result := tx.Where("id = ?", id).Delete(model.SVIndividualValues{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

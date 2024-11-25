package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVActualValuesRepository interface {
		Create(*gorm.DB, model.SVActualValues) error
		Update(*gorm.DB, model.SVActualValues) error
	}
	SVActualValuesRepository struct {
		Db *gorm.DB
	}
)

func NewSVActualValuesRepository(db *gorm.DB) ISVActualValuesRepository {
	return &SVActualValuesRepository{
		Db: db,
	}
}

func (s *SVActualValuesRepository) Create(tx *gorm.DB, model model.SVActualValues) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVActualValuesRepository) Update(tx *gorm.DB, model model.SVActualValues) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

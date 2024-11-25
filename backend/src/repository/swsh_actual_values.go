package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShActualValuesRepository interface {
		Create(*gorm.DB, model.SwShActualValues) error
		Update(*gorm.DB, model.SwShActualValues) error
	}
	SwShActualValuesRepository struct {
		Db *gorm.DB
	}
)

func NewSwShActualValuesRepository(db *gorm.DB) ISwShActualValuesRepository {
	return &SwShActualValuesRepository{
		Db: db,
	}
}

func (s *SwShActualValuesRepository) Create(tx *gorm.DB, model model.SwShActualValues) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SwShActualValuesRepository) Update(tx *gorm.DB, model model.SwShActualValues) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

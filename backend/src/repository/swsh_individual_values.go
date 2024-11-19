package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShIndividualValuesRepository interface {
		Create(*gorm.DB, model.SwShIndividualValues) error
	}
	SwShIndividualValuesRepository struct {
		Db *gorm.DB
	}
)

func NewSwShIndividualValuesRepository(db *gorm.DB) ISwShIndividualValuesRepository {
	return &SwShIndividualValuesRepository{
		Db: db,
	}
}

func (s *SwShIndividualValuesRepository) Create(tx *gorm.DB, model model.SwShIndividualValues) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShBasePointsRepository interface {
		Create(*gorm.DB, model.SwShBasePoints) error
	}
	SwShBasePointsRepository struct {
		Db *gorm.DB
	}
)

func NewSwShBasePointsRepository(db *gorm.DB) ISwShBasePointsRepository {
	return &SwShBasePointsRepository{
		Db: db,
	}
}

func (s *SwShBasePointsRepository) Create(tx *gorm.DB, model model.SwShBasePoints) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

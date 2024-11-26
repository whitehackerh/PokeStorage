package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShBasePointsRepository interface {
		Create(*gorm.DB, model.SwShBasePoints) error
		Update(*gorm.DB, model.SwShBasePoints) error
		Delete(*gorm.DB, string) error
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

func (s *SwShBasePointsRepository) Update(tx *gorm.DB, model model.SwShBasePoints) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SwShBasePointsRepository) Delete(tx *gorm.DB, id string) error {
	result := tx.Where("id = ?", id).Delete(model.SwShBasePoints{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

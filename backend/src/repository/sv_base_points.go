package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVBasePointsRepository interface {
		Create(*gorm.DB, model.SVBasePoints) error
		Update(*gorm.DB, model.SVBasePoints) error
		Delete(*gorm.DB, string) error
	}
	SVBasePointsRepository struct {
		Db *gorm.DB
	}
)

func NewSVBasePointsRepository(db *gorm.DB) ISVBasePointsRepository {
	return &SVBasePointsRepository{
		Db: db,
	}
}

func (s *SVBasePointsRepository) Create(tx *gorm.DB, model model.SVBasePoints) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVBasePointsRepository) Update(tx *gorm.DB, model model.SVBasePoints) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVBasePointsRepository) Delete(tx *gorm.DB, id string) error {
	result := tx.Where("id = ?", id).Delete(model.SVBasePoints{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

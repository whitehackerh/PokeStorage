package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	IGenderRepository interface {
		Fetch() ([]model.Gender, error)
	}
	GenderRepository struct {
		Db *gorm.DB
	}
)

func NewGenderRepository(db *gorm.DB) IGenderRepository {
	return &GenderRepository{
		Db: db,
	}
}

func (gr *GenderRepository) Fetch() ([]model.Gender, error) {
	var genders []model.Gender
	if result := gr.Db.Order("id ASC").Find(&genders); result.Error != nil {
		return []model.Gender{}, result.Error
	}
	return genders, nil
}

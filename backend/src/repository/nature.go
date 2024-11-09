package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	INatureRepository interface {
		Fetch() ([]model.Nature, error)
	}
	NatureRepository struct {
		Db *gorm.DB
	}
)

func NewNatureRepository(db *gorm.DB) INatureRepository {
	return &NatureRepository{
		Db: db,
	}
}

func (tr *NatureRepository) Fetch() ([]model.Nature, error) {
	var natures []model.Nature
	if result := tr.Db.Order("id ASC").Find(&natures); result.Error != nil {
		return []model.Nature{}, result.Error
	}
	return natures, nil
}

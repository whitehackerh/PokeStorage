package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVItemRepository interface {
		Fetch() ([]model.SVItem, error)
	}
	SVItemRepository struct {
		Db *gorm.DB
	}
)

func NewSVItemRepository(db *gorm.DB) ISVItemRepository {
	return &SVItemRepository{
		Db: db,
	}
}

func (r *SVItemRepository) Fetch() ([]model.SVItem, error) {
	var svItems []model.SVItem
	if result := r.Db.Order("id ASC").Find(&svItems); result.Error != nil {
		return []model.SVItem{}, result.Error
	}
	return svItems, nil
}

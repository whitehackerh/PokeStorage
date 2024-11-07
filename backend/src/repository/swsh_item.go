package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShItemRepository interface {
		Fetch() ([]model.SwShItem, error)
	}
	SwShItemRepository struct {
		Db *gorm.DB
	}
)

func NewSwShItemRepository(db *gorm.DB) ISwShItemRepository {
	return &SwShItemRepository{
		Db: db,
	}
}

func (r *SwShItemRepository) Fetch() ([]model.SwShItem, error) {
	var swShItems []model.SwShItem
	if result := r.Db.Order("id ASC").Find(&swShItems); result.Error != nil {
		return []model.SwShItem{}, result.Error
	}
	return swShItems, nil
}

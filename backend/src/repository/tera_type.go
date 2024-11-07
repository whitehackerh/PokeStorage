package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ITeraTypeRepository interface {
		Fetch() ([]model.TeraType, error)
	}
	TeraTypeRepository struct {
		Db *gorm.DB
	}
)

func NewTeraTypeRepository(db *gorm.DB) ITeraTypeRepository {
	return &TeraTypeRepository{
		Db: db,
	}
}

func (tr *TeraTypeRepository) Fetch() ([]model.TeraType, error) {
	var teraTypes []model.TeraType
	if result := tr.Db.Order("id ASC").Find(&teraTypes); result.Error != nil {
		return []model.TeraType{}, result.Error
	}
	return teraTypes, nil
}

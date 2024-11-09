package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShMoveRepository interface {
		Fetch() ([]model.SwShMoveRelation, error)
	}
	SwShMoveRepository struct {
		Db *gorm.DB
	}
)

func NewSwShMoveRepository(db *gorm.DB) ISwShMoveRepository {
	return &SwShMoveRepository{
		Db: db,
	}
}

func (r *SwShMoveRepository) Fetch() ([]model.SwShMoveRelation, error) {
	var swShMoveRelations []model.SwShMoveRelation
	if err := r.Db.Preload("Type").
		Preload("MoveCategory").
		Order("id asc").
		Find(&swShMoveRelations).Error; err != nil {
		return nil, err
	}
	return swShMoveRelations, nil
}

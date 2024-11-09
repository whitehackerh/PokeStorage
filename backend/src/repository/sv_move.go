package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVMoveRepository interface {
		Fetch() ([]model.SVMoveRelation, error)
	}
	SVMoveRepository struct {
		Db *gorm.DB
	}
)

func NewSVMoveRepository(db *gorm.DB) ISVMoveRepository {
	return &SVMoveRepository{
		Db: db,
	}
}

func (r *SVMoveRepository) Fetch() ([]model.SVMoveRelation, error) {
	var svMoveRelations []model.SVMoveRelation
	if err := r.Db.Preload("Type").
		Preload("MoveCategory").
		Order("id asc").
		Find(&svMoveRelations).Error; err != nil {
		return nil, err
	}
	return svMoveRelations, nil
}

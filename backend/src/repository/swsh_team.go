package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShTeamRepository interface {
		Create(*gorm.DB, model.SwShTeam) error
	}
	SwShTeamRepository struct {
		Db *gorm.DB
	}
)

func NewSwShTeamRepository(db *gorm.DB) ISwShTeamRepository {
	return &SwShTeamRepository{
		Db: db,
	}
}

func (s *SwShTeamRepository) Create(tx *gorm.DB, model model.SwShTeam) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVTeamRepository interface {
		Create(*gorm.DB, model.SVTeam) error
	}
	SVTeamRepository struct {
		Db *gorm.DB
	}
)

func NewSVTeamRepository(db *gorm.DB) ISVTeamRepository {
	return &SVTeamRepository{
		Db: db,
	}
}

func (s *SVTeamRepository) Create(tx *gorm.DB, model model.SVTeam) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

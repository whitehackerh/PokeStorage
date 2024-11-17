package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVBredPokemonRepository interface {
		Create(*gorm.DB, model.SVBredPokemon) error
	}
	SVBredPokemonRepository struct {
		Db *gorm.DB
	}
)

func NewSVBredPokemonRepository(db *gorm.DB) ISVBredPokemonRepository {
	return &SVBredPokemonRepository{
		Db: db,
	}
}

func (s *SVBredPokemonRepository) Create(tx *gorm.DB, model model.SVBredPokemon) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

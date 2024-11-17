package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShBredPokemonRepository interface {
		Create(*gorm.DB, model.SwShBredPokemon) error
	}
	SwShBredPokemonRepository struct {
		Db *gorm.DB
	}
)

func NewSwShBredPokemonRepository(db *gorm.DB) ISwShBredPokemonRepository {
	return &SwShBredPokemonRepository{
		Db: db,
	}
}

func (s *SwShBredPokemonRepository) Create(tx *gorm.DB, model model.SwShBredPokemon) error {
	result := tx.Create(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShPokemonRepository interface {
		Fetch() ([]model.SwShPokemonRelation, error)
	}
	SwShPokemonRepository struct {
		Db *gorm.DB
	}
)

func NewSwShPokemonRepository(db *gorm.DB) ISwShPokemonRepository {
	return &SwShPokemonRepository{
		Db: db,
	}
}

func (r *SwShPokemonRepository) Fetch() ([]model.SwShPokemonRelation, error) {
	var swShPokemonRelations []model.SwShPokemonRelation
	if err := r.Db.Preload("Type1").
		Preload("Type2").
		Preload("Ability1").
		Preload("Ability2").
		Preload("HiddenAbility").
		Preload("BaseStats").
		Preload("PresetHeldItem").
		Order("national_pokedex_no asc").
		Order("forme_no asc").
		Find(&swShPokemonRelations).Error; err != nil {
		return nil, err
	}
	return swShPokemonRelations, nil
}

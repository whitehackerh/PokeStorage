package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVPokemonRepository interface {
		Fetch() ([]model.SVPokemonRelation, error)
	}
	SVPokemonRepository struct {
		Db *gorm.DB
	}
)

func NewSVPokemonRepository(db *gorm.DB) ISVPokemonRepository {
	return &SVPokemonRepository{
		Db: db,
	}
}

func (r *SVPokemonRepository) Fetch() ([]model.SVPokemonRelation, error) {
	var svPokemonRelations []model.SVPokemonRelation
	if err := r.Db.Preload("Type1").
		Preload("Type2").
		Preload("Ability1").
		Preload("Ability2").
		Preload("HiddenAbility").
		Preload("BaseStats").
		Preload("PresetHeldItem").
		Order("national_pokedex_no asc").
		Order("forme_no asc").
		Find(&svPokemonRelations).Error; err != nil {
		return nil, err
	}
	return svPokemonRelations, nil
}

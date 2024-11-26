package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShBredPokemonRepository interface {
		Create(*gorm.DB, model.SwShBredPokemon) error
		FetchByUserId(string) ([]model.SwShBredPokemonRelation, error)
		Update(*gorm.DB, model.SwShBredPokemon) error
		FindById(string) (model.SwShBredPokemon, error)
		Delete(*gorm.DB, string) error
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

func (s *SwShBredPokemonRepository) FetchByUserId(userId string) ([]model.SwShBredPokemonRelation, error) {
	var swShBredPokemonRelations []model.SwShBredPokemonRelation
	if err := s.Db.Preload("SwShPokemonRelation").
		Preload("SwShPokemonRelation.Type1").
		Preload("SwShPokemonRelation.Type2").
		Preload("SwShPokemonRelation.Ability1").
		Preload("SwShPokemonRelation.Ability2").
		Preload("SwShPokemonRelation.HiddenAbility").
		Preload("SwShPokemonRelation.BaseStats").
		Preload("SwShPokemonRelation.PresetHeldItem").
		Preload("Gender").
		Preload("Ability").
		Preload("Nature").
		Preload("HeldItem").
		Preload("IndividualValues").
		Preload("BasePoints").
		Preload("ActualValues").
		Preload("Move1Relation").
		Preload("Move1Relation.Type").
		Preload("Move1Relation.MoveCategory").
		Preload("Move2Relation").
		Preload("Move2Relation.Type").
		Preload("Move2Relation.MoveCategory").
		Preload("Move3Relation").
		Preload("Move3Relation.Type").
		Preload("Move3Relation.MoveCategory").
		Preload("Move4Relation").
		Preload("Move4Relation.Type").
		Preload("Move4Relation.MoveCategory").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Order("created_at desc").
		Find(&swShBredPokemonRelations).
		Error; err != nil {
		return nil, err
	}
	return swShBredPokemonRelations, nil
}

func (s *SwShBredPokemonRepository) Update(tx *gorm.DB, model model.SwShBredPokemon) error {
	result := tx.Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SwShBredPokemonRepository) FindById(id string) (model.SwShBredPokemon, error) {
	var bredPokemon model.SwShBredPokemon
	result := s.Db.Where("id = ?", id).Find(&bredPokemon)
	if result.Error != nil {
		return model.SwShBredPokemon{}, result.Error
	}
	return bredPokemon, result.Error
}

func (s *SwShBredPokemonRepository) Delete(tx *gorm.DB, id string) error {
	result := tx.Where("id = ?", id).Delete(model.SwShBredPokemon{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

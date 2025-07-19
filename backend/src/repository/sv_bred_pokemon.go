package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVBredPokemonRepository interface {
		Create(*gorm.DB, model.SVBredPokemon) error
		FetchByUserId(string) ([]model.SVBredPokemonRelation, error)
		Update(*gorm.DB, model.SVBredPokemon) error
		FindById(string) (model.SVBredPokemon, error)
		Delete(*gorm.DB, string) error
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

func (s *SVBredPokemonRepository) FetchByUserId(userId string) ([]model.SVBredPokemonRelation, error) {
	var svBredPokemonRelations []model.SVBredPokemonRelation
	if err := s.Db.Preload("SVPokemonRelation").
		Preload("SVPokemonRelation.Type1").
		Preload("SVPokemonRelation.Type2").
		Preload("SVPokemonRelation.Ability1").
		Preload("SVPokemonRelation.Ability2").
		Preload("SVPokemonRelation.HiddenAbility").
		Preload("SVPokemonRelation.BaseStats").
		Preload("SVPokemonRelation.PresetHeldItem").
		Preload("Gender").
		Preload("TeraType").
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
		Find(&svBredPokemonRelations).
		Error; err != nil {
		return nil, err
	}
	return svBredPokemonRelations, nil
}

func (s *SVBredPokemonRepository) Update(tx *gorm.DB, model model.SVBredPokemon) error {
	result := tx.Select("*").Omit("CreatedAt", "DeletedAt").Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SVBredPokemonRepository) FindById(id string) (model.SVBredPokemon, error) {
	var bredPokemon model.SVBredPokemon
	result := s.Db.Where("id = ?", id).Find(&bredPokemon)
	if result.Error != nil {
		return model.SVBredPokemon{}, result.Error
	}
	return bredPokemon, result.Error
}

func (s *SVBredPokemonRepository) Delete(tx *gorm.DB, id string) error {
	result := tx.Where("id = ?", id).Delete(model.SVBredPokemon{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

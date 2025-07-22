package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVTeamRepository interface {
		Create(*gorm.DB, model.SVTeam) error
		ExistsBredPokemon(string) (bool, error)
		FetchByUserId(string) ([]model.SVTeamRelation, error)
		Update(*gorm.DB, model.SVTeam) error
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

func (s *SVTeamRepository) ExistsBredPokemon(bredPokemonId string) (bool, error) {
	var count int64
	result := s.Db.Model(&model.SVTeam{}).
		Where("deleted_at IS NULL AND (bred_pokemon_1_id = ? OR bred_pokemon_2_id = ? OR bred_pokemon_3_id = ? OR bred_pokemon_4_id = ? OR bred_pokemon_5_id = ? OR bred_pokemon_6_id = ?)",
			bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (s *SVTeamRepository) FetchByUserId(userId string) ([]model.SVTeamRelation, error) {
	var svTeams []model.SVTeamRelation
	if err := s.Db.
		Preload("BredPokemon1").
		Preload("BredPokemon1.SVPokemonRelation").
		Preload("BredPokemon1.SVPokemonRelation.Type1").
		Preload("BredPokemon1.SVPokemonRelation.Type2").
		Preload("BredPokemon1.SVPokemonRelation.Ability1").
		Preload("BredPokemon1.SVPokemonRelation.Ability2").
		Preload("BredPokemon1.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon1.SVPokemonRelation.BaseStats").
		Preload("BredPokemon1.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon1.Gender").
		Preload("BredPokemon1.TeraType").
		Preload("BredPokemon1.Ability").
		Preload("BredPokemon1.Nature").
		Preload("BredPokemon1.HeldItem").
		Preload("BredPokemon1.IndividualValues").
		Preload("BredPokemon1.BasePoints").
		Preload("BredPokemon1.ActualValues").
		Preload("BredPokemon1.Move1Relation").
		Preload("BredPokemon1.Move1Relation.Type").
		Preload("BredPokemon1.Move1Relation.MoveCategory").
		Preload("BredPokemon1.Move2Relation").
		Preload("BredPokemon1.Move2Relation.Type").
		Preload("BredPokemon1.Move2Relation.MoveCategory").
		Preload("BredPokemon1.Move3Relation").
		Preload("BredPokemon1.Move3Relation.Type").
		Preload("BredPokemon1.Move3Relation.MoveCategory").
		Preload("BredPokemon1.Move4Relation").
		Preload("BredPokemon1.Move4Relation.Type").
		Preload("BredPokemon1.Move4Relation.MoveCategory").
		Preload("BredPokemon2").
		Preload("BredPokemon2.SVPokemonRelation").
		Preload("BredPokemon2.SVPokemonRelation.Type1").
		Preload("BredPokemon2.SVPokemonRelation.Type2").
		Preload("BredPokemon2.SVPokemonRelation.Ability1").
		Preload("BredPokemon2.SVPokemonRelation.Ability2").
		Preload("BredPokemon2.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon2.SVPokemonRelation.BaseStats").
		Preload("BredPokemon2.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon2.Gender").
		Preload("BredPokemon2.TeraType").
		Preload("BredPokemon2.Ability").
		Preload("BredPokemon2.Nature").
		Preload("BredPokemon2.HeldItem").
		Preload("BredPokemon2.IndividualValues").
		Preload("BredPokemon2.BasePoints").
		Preload("BredPokemon2.ActualValues").
		Preload("BredPokemon2.Move1Relation").
		Preload("BredPokemon2.Move1Relation.Type").
		Preload("BredPokemon2.Move1Relation.MoveCategory").
		Preload("BredPokemon2.Move2Relation").
		Preload("BredPokemon2.Move2Relation.Type").
		Preload("BredPokemon2.Move2Relation.MoveCategory").
		Preload("BredPokemon2.Move3Relation").
		Preload("BredPokemon2.Move3Relation.Type").
		Preload("BredPokemon2.Move3Relation.MoveCategory").
		Preload("BredPokemon2.Move4Relation").
		Preload("BredPokemon2.Move4Relation.Type").
		Preload("BredPokemon2.Move4Relation.MoveCategory").
		Preload("BredPokemon3").
		Preload("BredPokemon3.SVPokemonRelation").
		Preload("BredPokemon3.SVPokemonRelation.Type1").
		Preload("BredPokemon3.SVPokemonRelation.Type2").
		Preload("BredPokemon3.SVPokemonRelation.Ability1").
		Preload("BredPokemon3.SVPokemonRelation.Ability2").
		Preload("BredPokemon3.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon3.SVPokemonRelation.BaseStats").
		Preload("BredPokemon3.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon3.Gender").
		Preload("BredPokemon3.TeraType").
		Preload("BredPokemon3.Ability").
		Preload("BredPokemon3.Nature").
		Preload("BredPokemon3.HeldItem").
		Preload("BredPokemon3.IndividualValues").
		Preload("BredPokemon3.BasePoints").
		Preload("BredPokemon3.ActualValues").
		Preload("BredPokemon3.Move1Relation").
		Preload("BredPokemon3.Move1Relation.Type").
		Preload("BredPokemon3.Move1Relation.MoveCategory").
		Preload("BredPokemon3.Move2Relation").
		Preload("BredPokemon3.Move2Relation.Type").
		Preload("BredPokemon3.Move2Relation.MoveCategory").
		Preload("BredPokemon3.Move3Relation").
		Preload("BredPokemon3.Move3Relation.Type").
		Preload("BredPokemon3.Move3Relation.MoveCategory").
		Preload("BredPokemon3.Move4Relation").
		Preload("BredPokemon3.Move4Relation.Type").
		Preload("BredPokemon3.Move4Relation.MoveCategory").
		Preload("BredPokemon4").
		Preload("BredPokemon4.SVPokemonRelation").
		Preload("BredPokemon4.SVPokemonRelation.Type1").
		Preload("BredPokemon4.SVPokemonRelation.Type2").
		Preload("BredPokemon4.SVPokemonRelation.Ability1").
		Preload("BredPokemon4.SVPokemonRelation.Ability2").
		Preload("BredPokemon4.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon4.SVPokemonRelation.BaseStats").
		Preload("BredPokemon4.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon4.Gender").
		Preload("BredPokemon4.TeraType").
		Preload("BredPokemon4.Ability").
		Preload("BredPokemon4.Nature").
		Preload("BredPokemon4.HeldItem").
		Preload("BredPokemon4.IndividualValues").
		Preload("BredPokemon4.BasePoints").
		Preload("BredPokemon4.ActualValues").
		Preload("BredPokemon4.Move1Relation").
		Preload("BredPokemon4.Move1Relation.Type").
		Preload("BredPokemon4.Move1Relation.MoveCategory").
		Preload("BredPokemon4.Move2Relation").
		Preload("BredPokemon4.Move2Relation.Type").
		Preload("BredPokemon4.Move2Relation.MoveCategory").
		Preload("BredPokemon4.Move3Relation").
		Preload("BredPokemon4.Move3Relation.Type").
		Preload("BredPokemon4.Move3Relation.MoveCategory").
		Preload("BredPokemon4.Move4Relation").
		Preload("BredPokemon4.Move4Relation.Type").
		Preload("BredPokemon4.Move4Relation.MoveCategory").
		Preload("BredPokemon5").
		Preload("BredPokemon5.SVPokemonRelation").
		Preload("BredPokemon5.SVPokemonRelation.Type1").
		Preload("BredPokemon5.SVPokemonRelation.Type2").
		Preload("BredPokemon5.SVPokemonRelation.Ability1").
		Preload("BredPokemon5.SVPokemonRelation.Ability2").
		Preload("BredPokemon5.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon5.SVPokemonRelation.BaseStats").
		Preload("BredPokemon5.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon5.Gender").
		Preload("BredPokemon5.TeraType").
		Preload("BredPokemon5.Ability").
		Preload("BredPokemon5.Nature").
		Preload("BredPokemon5.HeldItem").
		Preload("BredPokemon5.IndividualValues").
		Preload("BredPokemon5.BasePoints").
		Preload("BredPokemon5.ActualValues").
		Preload("BredPokemon5.Move1Relation").
		Preload("BredPokemon5.Move1Relation.Type").
		Preload("BredPokemon5.Move1Relation.MoveCategory").
		Preload("BredPokemon5.Move2Relation").
		Preload("BredPokemon5.Move2Relation.Type").
		Preload("BredPokemon5.Move2Relation.MoveCategory").
		Preload("BredPokemon5.Move3Relation").
		Preload("BredPokemon5.Move3Relation.Type").
		Preload("BredPokemon5.Move3Relation.MoveCategory").
		Preload("BredPokemon5.Move4Relation").
		Preload("BredPokemon5.Move4Relation.Type").
		Preload("BredPokemon5.Move4Relation.MoveCategory").
		Preload("BredPokemon6").
		Preload("BredPokemon6.SVPokemonRelation").
		Preload("BredPokemon6.SVPokemonRelation.Type1").
		Preload("BredPokemon6.SVPokemonRelation.Type2").
		Preload("BredPokemon6.SVPokemonRelation.Ability1").
		Preload("BredPokemon6.SVPokemonRelation.Ability2").
		Preload("BredPokemon6.SVPokemonRelation.HiddenAbility").
		Preload("BredPokemon6.SVPokemonRelation.BaseStats").
		Preload("BredPokemon6.SVPokemonRelation.PresetHeldItem").
		Preload("BredPokemon6.Gender").
		Preload("BredPokemon6.TeraType").
		Preload("BredPokemon6.Ability").
		Preload("BredPokemon6.Nature").
		Preload("BredPokemon6.HeldItem").
		Preload("BredPokemon6.IndividualValues").
		Preload("BredPokemon6.BasePoints").
		Preload("BredPokemon6.ActualValues").
		Preload("BredPokemon6.Move1Relation").
		Preload("BredPokemon6.Move1Relation.Type").
		Preload("BredPokemon6.Move1Relation.MoveCategory").
		Preload("BredPokemon6.Move2Relation").
		Preload("BredPokemon6.Move2Relation.Type").
		Preload("BredPokemon6.Move2Relation.MoveCategory").
		Preload("BredPokemon6.Move3Relation").
		Preload("BredPokemon6.Move3Relation.Type").
		Preload("BredPokemon6.Move3Relation.MoveCategory").
		Preload("BredPokemon6.Move4Relation").
		Preload("BredPokemon6.Move4Relation.Type").
		Preload("BredPokemon6.Move4Relation.MoveCategory").
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL").
		Order("created_at desc").
		Find(&svTeams).Error; err != nil {
		return nil, err
	}
	return svTeams, nil
}

func (s *SVTeamRepository) Update(tx *gorm.DB, model model.SVTeam) error {
	result := tx.Select("*").Omit("CreatedAt", "DeletedAt").Updates(&model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

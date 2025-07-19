package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISwShTeamRepository interface {
		Create(*gorm.DB, model.SwShTeam) error
		ExistsBredPokemon(string) (bool, error)
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

func (s *SwShTeamRepository) ExistsBredPokemon(bredPokemonId string) (bool, error) {
	var count int64
	result := s.Db.Model(&model.SwShTeam{}).
		Where("deleted_at IS NULL AND (bred_pokemon_1_id = ? OR bred_pokemon_2_id = ? OR bred_pokemon_3_id = ? OR bred_pokemon_4_id = ? OR bred_pokemon_5_id = ? OR bred_pokemon_6_id = ?)",
			bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId, bredPokemonId).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

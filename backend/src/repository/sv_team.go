package repository

import (
	"github.com/whitehackerh/PokeStorage/src/model"
	"gorm.io/gorm"
)

type (
	ISVTeamRepository interface {
		Create(*gorm.DB, model.SVTeam) error
		ExistsBredPokemon(string) (bool, error)
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

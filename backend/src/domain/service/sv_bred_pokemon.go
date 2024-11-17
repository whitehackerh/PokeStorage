package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
)

type (
	ISVBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.SVBredPokemon) entity.SVBredPokemon
	}
	SVBredPokemonService struct {
		bredPokemonService IBredPokemonService
	}
)

func NewSVBredPokemonService(bredPokemonService IBredPokemonService) ISVBredPokemonService {
	return &SVBredPokemonService{
		bredPokemonService: bredPokemonService,
	}
}

func (s *SVBredPokemonService) MakeEntityFromApiSchema(schema api_schema.SVBredPokemon) entity.SVBredPokemon {
	return entity.NewSVBredPokemon(
		s.bredPokemonService.MakeEntityFromApiSchema(schema.BredPokemon),
		entity.NewTeraType(
			schema.TeraType.Id,
			schema.TeraType.Name,
		),
	)
}

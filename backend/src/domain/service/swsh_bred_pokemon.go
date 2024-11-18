package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
)

type (
	ISwShBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.SwShBredPokemon, string) entity.SwShBredPokemon
	}
	SwShBredPokemonService struct {
		bredPokemonService IBredPokemonService
	}
)

func NewSwShBredPokemonService(bredPokemonService IBredPokemonService) ISwShBredPokemonService {
	return &SwShBredPokemonService{
		bredPokemonService: bredPokemonService,
	}
}

func (s *SwShBredPokemonService) MakeEntityFromApiSchema(schema api_schema.SwShBredPokemon, userId string) entity.SwShBredPokemon {
	return entity.NewSwShBredPokemon(
		s.bredPokemonService.MakeEntityFromApiSchema(schema.BredPokemon, userId),
		schema.IsGigantamax,
	)
}

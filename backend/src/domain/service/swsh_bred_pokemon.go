package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
)

type (
	ISwShBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.SwShBredPokemon) entity.SwShBredPokemon
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

func (s *SwShBredPokemonService) MakeEntityFromApiSchema(schema api_schema.SwShBredPokemon) entity.SwShBredPokemon {
	return entity.NewSwShBredPokemon(
		s.bredPokemonService.MakeEntityFromApiSchema(schema.BredPokemon),
		schema.IsGigantamax,
	)
}

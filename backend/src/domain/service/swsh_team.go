package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/enum"
)

type (
	ISwShTeamService interface {
		MakeEntityFromApiSchema(api_schema.PostPutTeam, string) entity.SwShTeam
	}
	SwShTeamService struct {
		teamService ITeamService
	}
)

func NewSwShTeamService(teamService ITeamService) ISwShTeamService {
	return &SwShTeamService{
		teamService: teamService,
	}
}

func (s *SwShTeamService) MakeEntityFromApiSchema(schema api_schema.PostPutTeam, userId string) entity.SwShTeam {
	var bredPokemons = make([]*entity.SwShBredPokemon, enum.MaxTeamSize)
	for i, bredPokemonId := range schema.BredPokemonIds {
		if bredPokemonId == nil {
			bredPokemons[i] = nil
		} else {
			bredPokemon := entity.NewSwShBredPokemon(
				entity.NewBredPokemon(
					*bredPokemonId,
					"",
					0,
					0,
					0,
					"",
					nil,
					entity.Gender{},
					0,
					[]entity.Type{},
					entity.Ability{},
					entity.Nature{},
					nil,
					entity.BaseStats{},
					entity.IndividualValues{},
					entity.BasePoints{},
					entity.ActualValues{},
					[]entity.Move{},
					nil,
				),
				false,
			)
			bredPokemons[i] = &bredPokemon
		}
	}
	return entity.NewSwShTeam(
		s.teamService.MakeEntityFromApiSchema(schema, userId),
		bredPokemons,
	)
}

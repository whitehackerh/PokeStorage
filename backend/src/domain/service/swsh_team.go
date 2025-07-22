package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/enum"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShTeamService interface {
		MakeEntityFromApiSchema(api_schema.PostPutTeam, string) entity.SwShTeam
		MakeEntityFromModel(model.SwShTeamRelation) entity.SwShTeam
	}
	SwShTeamService struct {
		teamService            ITeamService
		swshBredPokemonService ISwShBredPokemonService
	}
)

func NewSwShTeamService(teamService ITeamService, swshBredPokemonService ISwShBredPokemonService) ISwShTeamService {
	return &SwShTeamService{
		teamService:            teamService,
		swshBredPokemonService: swshBredPokemonService,
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

func (s *SwShTeamService) MakeEntityFromModel(model model.SwShTeamRelation) entity.SwShTeam {
	bredPokemons := make([]*entity.SwShBredPokemon, enum.MaxTeamSize)
	if model.BredPokemon1 == nil {
		bredPokemons[0] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon1)
		bredPokemons[0] = &bredPokemon
	}
	if model.BredPokemon2 == nil {
		bredPokemons[1] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon2)
		bredPokemons[1] = &bredPokemon
	}
	if model.BredPokemon3 == nil {
		bredPokemons[2] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon3)
		bredPokemons[2] = &bredPokemon
	}
	if model.BredPokemon4 == nil {
		bredPokemons[3] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon4)
		bredPokemons[3] = &bredPokemon
	}
	if model.BredPokemon5 == nil {
		bredPokemons[4] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon5)
		bredPokemons[4] = &bredPokemon
	}
	if model.BredPokemon6 == nil {
		bredPokemons[5] = nil
	} else {
		bredPokemon := s.swshBredPokemonService.MakeEntityFromModel(*model.BredPokemon6)
		bredPokemons[5] = &bredPokemon
	}

	return entity.NewSwShTeam(
		entity.NewTeam(
			model.Id,
			model.UserId,
			model.Name,
			model.Note,
		),
		bredPokemons,
	)
}

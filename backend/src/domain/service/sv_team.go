package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/enum"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVTeamService interface {
		MakeEntityFromApiSchema(api_schema.PostPutTeam, string) entity.SVTeam
		MakeEntityFromModel(model.SVTeamRelation) entity.SVTeam
	}
	SVTeamService struct {
		teamService          ITeamService
		svBredPokemonService ISVBredPokemonService
	}
)

func NewSVTeamService(teamService ITeamService, svBredPokemonService ISVBredPokemonService) ISVTeamService {
	return &SVTeamService{
		teamService:          teamService,
		svBredPokemonService: svBredPokemonService,
	}
}

func (s *SVTeamService) MakeEntityFromApiSchema(schema api_schema.PostPutTeam, userId string) entity.SVTeam {
	var bredPokemons = make([]*entity.SVBredPokemon, enum.MaxTeamSize)
	for i, bredPokemonId := range schema.BredPokemonIds {
		if bredPokemonId == nil {
			bredPokemons[i] = nil
		} else {
			bredPokemon := entity.NewSVBredPokemon(
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
				entity.TeraType{},
			)
			bredPokemons[i] = &bredPokemon
		}
	}
	return entity.NewSVTeam(
		s.teamService.MakeEntityFromApiSchema(schema, userId),
		bredPokemons,
	)
}

func (s *SVTeamService) MakeEntityFromModel(model model.SVTeamRelation) entity.SVTeam {
	bredPokemons := make([]*entity.SVBredPokemon, enum.MaxTeamSize)
	if model.BredPokemon1 == nil {
		bredPokemons[0] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon1)
		bredPokemons[0] = &bredPokemon
	}
	if model.BredPokemon2 == nil {
		bredPokemons[1] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon2)
		bredPokemons[1] = &bredPokemon
	}
	if model.BredPokemon3 == nil {
		bredPokemons[2] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon3)
		bredPokemons[2] = &bredPokemon
	}
	if model.BredPokemon4 == nil {
		bredPokemons[3] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon4)
		bredPokemons[3] = &bredPokemon
	}
	if model.BredPokemon5 == nil {
		bredPokemons[4] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon5)
		bredPokemons[4] = &bredPokemon
	}
	if model.BredPokemon6 == nil {
		bredPokemons[5] = nil
	} else {
		bredPokemon := s.svBredPokemonService.MakeEntityFromModel(*model.BredPokemon6)
		bredPokemons[5] = &bredPokemon
	}

	return entity.NewSVTeam(
		entity.NewTeam(
			model.Id,
			model.UserId,
			model.Name,
			model.Note,
		),
		bredPokemons,
	)
}

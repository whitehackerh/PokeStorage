package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
)

type (
	ISVTeamService interface {
		MakeEntityFromApiSchema(api_schema.PostPutTeam, string) entity.SVTeam
	}
	SVTeamService struct {
		teamService ITeamService
	}
)

func NewSVTeamService(teamService ITeamService) ISVTeamService {
	return &SVTeamService{
		teamService: teamService,
	}
}

func (s *SVTeamService) MakeEntityFromApiSchema(schema api_schema.PostPutTeam, userId string) entity.SVTeam {
	var bredPokemons []entity.SVBredPokemon
	for _, bredPokemonId := range schema.BredPokemonIds {
		bredPokemons = append(bredPokemons, entity.NewSVBredPokemon(
			entity.NewBredPokemon(
				bredPokemonId,
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
		))
	}
	return entity.NewSVTeam(
		s.teamService.MakeEntityFromApiSchema(schema, userId),
		bredPokemons,
	)
}

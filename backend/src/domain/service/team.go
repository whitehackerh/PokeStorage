package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/util"
)

type (
	ITeamService interface {
		MakeEntityFromApiSchema(api_schema.PostPutTeam, string) entity.Team
	}
	TeamService struct{}
)

func NewTeamService() ITeamService {
	return &TeamService{}
}

func (s *TeamService) MakeEntityFromApiSchema(schema api_schema.PostPutTeam, userId string) entity.Team {
	var id string
	if schema.Id == "" {
		id = util.NewUUID()
	} else {
		id = schema.Id
	}

	return entity.NewTeam(
		id,
		userId,
		schema.Name,
		schema.Note,
	)
}

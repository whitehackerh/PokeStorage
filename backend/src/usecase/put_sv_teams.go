package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PutSVTeamsUseCase interface {
		Execute(input PutSVTeamsInput, userId string, tx *gorm.DB) (PutSVTeamsOutput, error)
	}
	PutSVTeamsInput struct {
		Team api_schema.PostPutTeam `json:"team"`
	}
	PutSVTeamsPresenter interface {
		Output() PutSVTeamsOutput
	}
	PutSVTeamsOutput     struct{}
	PutSVTeamsInteractor struct {
		service    service.ISVTeamService
		repository repository.ISVTeamRepository
		presenter  PutSVTeamsPresenter
	}
)

func NewPutSVTeamsInteractor(
	service service.ISVTeamService,
	repository repository.ISVTeamRepository,
	presenter PutSVTeamsPresenter,
) PutSVTeamsUseCase {
	return &PutSVTeamsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *PutSVTeamsInteractor) Execute(input PutSVTeamsInput, userId string, tx *gorm.DB) (PutSVTeamsOutput, error) {
	team := interactor.service.MakeEntityFromApiSchema(input.Team, userId)

	if err := interactor.repository.Update(
		tx,
		converter.SVTeamEntityToModel(team),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

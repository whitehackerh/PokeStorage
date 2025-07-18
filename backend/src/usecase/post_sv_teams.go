package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PostSVTeamsUseCase interface {
		Execute(input PostSVTeamsInput, userId string, tx *gorm.DB) (PostSVTeamsOutput, error)
	}
	PostSVTeamsInput struct {
		Team api_schema.PostPutTeam `json:"team"`
	}
	PostSVTeamsPresenter interface {
		Output() PostSVTeamsOutput
	}
	PostSVTeamsOutput     struct{}
	PostSVTeamsInteractor struct {
		service    service.ISVTeamService
		repository repository.ISVTeamRepository
		presenter  PostSVTeamsPresenter
	}
)

func NewPostSVTeamsInteractor(
	service service.ISVTeamService,
	repository repository.ISVTeamRepository,
	presenter PostSVTeamsPresenter,
) PostSVTeamsUseCase {
	return &PostSVTeamsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *PostSVTeamsInteractor) Execute(input PostSVTeamsInput, userId string, tx *gorm.DB) (PostSVTeamsOutput, error) {
	team := interactor.service.MakeEntityFromApiSchema(input.Team, userId)

	if err := interactor.repository.Create(
		tx,
		converter.SVTeamEntityToModel(team),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

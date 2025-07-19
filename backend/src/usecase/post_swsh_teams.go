package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PostSwShTeamsUseCase interface {
		Execute(input PostSwShTeamsInput, userId string, tx *gorm.DB) (PostSwShTeamsOutput, error)
	}
	PostSwShTeamsInput struct {
		Team api_schema.PostPutTeam `json:"team"`
	}
	PostSwShTeamsPresenter interface {
		Output() PostSwShTeamsOutput
	}
	PostSwShTeamsOutput     struct{}
	PostSwShTeamsInteractor struct {
		service    service.ISwShTeamService
		repository repository.ISwShTeamRepository
		presenter  PostSwShTeamsPresenter
	}
)

func NewPostSwShTeamsInteractor(
	service service.ISwShTeamService,
	repository repository.ISwShTeamRepository,
	presenter PostSwShTeamsPresenter,
) PostSwShTeamsUseCase {
	return &PostSwShTeamsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *PostSwShTeamsInteractor) Execute(input PostSwShTeamsInput, userId string, tx *gorm.DB) (PostSwShTeamsOutput, error) {
	team := interactor.service.MakeEntityFromApiSchema(input.Team, userId)

	if err := interactor.repository.Create(
		tx,
		converter.SwShTeamEntityToModel(team),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

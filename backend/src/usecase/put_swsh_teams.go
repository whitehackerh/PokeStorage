package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PutSwShTeamsUseCase interface {
		Execute(input PutSwShTeamsInput, userId string, tx *gorm.DB) (PutSwShTeamsOutput, error)
	}
	PutSwShTeamsInput struct {
		Team api_schema.PostPutTeam `json:"team"`
	}
	PutSwShTeamsPresenter interface {
		Output() PutSwShTeamsOutput
	}
	PutSwShTeamsOutput     struct{}
	PutSwShTeamsInteractor struct {
		service    service.ISwShTeamService
		repository repository.ISwShTeamRepository
		presenter  PutSwShTeamsPresenter
	}
)

func NewPutSwShTeamsInteractor(
	service service.ISwShTeamService,
	repository repository.ISwShTeamRepository,
	presenter PutSwShTeamsPresenter,
) PutSwShTeamsUseCase {
	return &PutSwShTeamsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *PutSwShTeamsInteractor) Execute(input PutSwShTeamsInput, userId string, tx *gorm.DB) (PutSwShTeamsOutput, error) {
	team := interactor.service.MakeEntityFromApiSchema(input.Team, userId)

	if err := interactor.repository.Update(
		tx,
		converter.SwShTeamEntityToModel(team),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

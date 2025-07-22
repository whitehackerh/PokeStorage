package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSwShTeamsUseCase interface {
		Execute(input GetSwShTeamsInput) (GetSwShTeamsOutput, error)
	}
	GetSwShTeamsInput struct {
		UserId string
	}
	GetSwShTeamsPresenter interface {
		Output([]entity.SwShTeam) GetSwShTeamsOutput
	}
	GetSwShTeamsOutput struct {
		Teams []api_schema.SwShTeam `json:"teams"`
	}
	GetSwShTeamsInteractor struct {
		service               service.ISwShTeamService
		teamRepository        repository.ISwShTeamRepository
		bredPokemonRepository repository.ISwShBredPokemonRepository
		presenter             GetSwShTeamsPresenter
	}
)

func NewGetSwShTeamsInteractor(
	service service.ISwShTeamService,
	teamRepository repository.ISwShTeamRepository,
	bredPokemonRepository repository.ISwShBredPokemonRepository,
	presenter GetSwShTeamsPresenter,
) GetSwShTeamsUseCase {
	return &GetSwShTeamsInteractor{
		service:               service,
		teamRepository:        teamRepository,
		bredPokemonRepository: bredPokemonRepository,
		presenter:             presenter,
	}
}

func (interactor *GetSwShTeamsInteractor) Execute(input GetSwShTeamsInput) (GetSwShTeamsOutput, error) {
	models, err := interactor.teamRepository.FetchByUserId(input.UserId)
	if err != nil {
		return interactor.presenter.Output([]entity.SwShTeam{}), err
	}

	var teams []entity.SwShTeam
	for _, model := range models {
		teams = append(teams, interactor.service.MakeEntityFromModel(model))
	}

	return interactor.presenter.Output(teams), nil
}

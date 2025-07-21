package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSVTeamsUseCase interface {
		Execute(input GetSVTeamsInput) (GetSVTeamsOutput, error)
	}
	GetSVTeamsInput struct {
		UserId string
	}
	GetSVTeamsPresenter interface {
		Output([]entity.SVTeam) GetSVTeamsOutput
	}
	GetSVTeamsOutput struct {
		Teams []api_schema.SVTeam `json:"teams"`
	}
	GetSVTeamsInteractor struct {
		service               service.ISVTeamService
		teamRepository        repository.ISVTeamRepository
		bredPokemonRepository repository.ISVBredPokemonRepository
		presenter             GetSVTeamsPresenter
	}
)

func NewGetSVTeamsInteractor(
	service service.ISVTeamService,
	teamRepository repository.ISVTeamRepository,
	bredPokemonRepository repository.ISVBredPokemonRepository,
	presenter GetSVTeamsPresenter,
) GetSVTeamsUseCase {
	return &GetSVTeamsInteractor{
		service:               service,
		teamRepository:        teamRepository,
		bredPokemonRepository: bredPokemonRepository,
		presenter:             presenter,
	}
}

func (interactor *GetSVTeamsInteractor) Execute(input GetSVTeamsInput) (GetSVTeamsOutput, error) {
	models, err := interactor.teamRepository.FetchByUserId(input.UserId)
	if err != nil {
		return interactor.presenter.Output([]entity.SVTeam{}), err
	}

	var teams []entity.SVTeam
	for _, model := range models {
		teams = append(teams, interactor.service.MakeEntityFromModel(model))
	}

	return interactor.presenter.Output(teams), nil
}

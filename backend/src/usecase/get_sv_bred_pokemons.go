package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSVBredPokemonsUseCase interface {
		Execute(input GetSVBredPokemonsInput) (GetSVBredPokemonsOutput, error)
	}
	GetSVBredPokemonsInput struct {
		UserId string
	}
	GetSVBredPokemonsPresenter interface {
		Output([]entity.SVBredPokemon) GetSVBredPokemonsOutput
	}
	GetSVBredPokemonsOutput struct {
		BredPokemons []api_schema.SVBredPokemon `json:"bred_pokemons"`
	}
	GetSVBredPokemonsInteractor struct {
		service    service.ISVBredPokemonService
		repository repository.ISVBredPokemonRepository
		presenter  GetSVBredPokemonsPresenter
	}
)

func NewGetSVBredPokemonsInteractor(
	service service.ISVBredPokemonService,
	repository repository.ISVBredPokemonRepository,
	presenter GetSVBredPokemonsPresenter,
) GetSVBredPokemonsUseCase {
	return &GetSVBredPokemonsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetSVBredPokemonsInteractor) Execute(input GetSVBredPokemonsInput) (GetSVBredPokemonsOutput, error) {
	models, err := interactor.repository.FetchByUserId(input.UserId)
	if err != nil {
		return interactor.presenter.Output([]entity.SVBredPokemon{}), err
	}

	var bredPokemons []entity.SVBredPokemon
	for _, model := range models {
		bredPokemons = append(bredPokemons, interactor.service.MakeEntityFromModel(model))
	}

	return interactor.presenter.Output(bredPokemons), nil
}

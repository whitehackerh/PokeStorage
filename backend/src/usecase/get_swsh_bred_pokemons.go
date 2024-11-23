package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSwShBredPokemonsUseCase interface {
		Execute(input GetSwShBredPokemonsInput) (GetSwShBredPokemonsOutput, error)
	}
	GetSwShBredPokemonsInput struct {
		UserId string
	}
	GetSwShBredPokemonsPresenter interface {
		Output([]entity.SwShBredPokemon) GetSwShBredPokemonsOutput
	}
	GetSwShBredPokemonsOutput struct {
		BredPokemons []api_schema.SwShBredPokemon `json:"bred_pokemons"`
	}
	GetSwShBredPokemonsInteractor struct {
		service    service.ISwShBredPokemonService
		repository repository.ISwShBredPokemonRepository
		presenter  GetSwShBredPokemonsPresenter
	}
)

func NewGetSwShBredPokemonsInteractor(
	service service.ISwShBredPokemonService,
	repository repository.ISwShBredPokemonRepository,
	presenter GetSwShBredPokemonsPresenter,
) GetSwShBredPokemonsUseCase {
	return &GetSwShBredPokemonsInteractor{
		service:    service,
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetSwShBredPokemonsInteractor) Execute(input GetSwShBredPokemonsInput) (GetSwShBredPokemonsOutput, error) {
	models, err := interactor.repository.FetchByUserId(input.UserId)
	if err != nil {
		return interactor.presenter.Output([]entity.SwShBredPokemon{}), err
	}

	var bredPokemons []entity.SwShBredPokemon
	for _, model := range models {
		bredPokemons = append(bredPokemons, interactor.service.MakeEntityFromModel(model))
	}

	return interactor.presenter.Output(bredPokemons), nil
}

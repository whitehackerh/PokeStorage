package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSwShPokemonsUseCase interface {
		Execute(input GetSwShPokemonsInput) (GetSwShPokemonsOutput, error)
	}
	GetSwShPokemonsInput     struct{}
	GetSwShPokemonsPresenter interface {
		Output([]entity.Pokemon) GetSwShPokemonsOutput
	}
	GetSwShPokemonsOutput struct {
		Pokemons []api_schema.Pokemon `json:"pokemons"`
	}
	GetSwShPokemonsInteractor struct {
		repository repository.ISwShPokemonRepository
		service    service.ISwShPokemonService
		presenter  GetSwShPokemonsPresenter
	}
)

func NewGetSwShPokemonsInteractor(
	repository repository.ISwShPokemonRepository,
	service service.ISwShPokemonService,
	presenter GetSwShPokemonsPresenter,
) GetSwShPokemonsUseCase {
	return &GetSwShPokemonsInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSwShPokemonsInteractor) Execute(input GetSwShPokemonsInput) (GetSwShPokemonsOutput, error) {
	pokemons, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Pokemon{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(pokemons)), err
}

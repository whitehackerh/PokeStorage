package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetPokemonsUseCase interface {
		Execute(input GetPokemonsInput) (GetPokemonsOutput, error)
	}
	GetPokemonsInput     struct{}
	GetPokemonsPresenter interface {
		Output([]entity.Pokemon) GetPokemonsOutput
	}
	GetPokemonsOutput struct {
		Pokemons []api_schema.Pokemon `json:"pokemons"`
	}
	GetSwShPokemonsInteractor struct {
		repository repository.ISwShPokemonRepository
		service    service.ISwShPokemonService
		presenter  GetPokemonsPresenter
	}
	GetSVPokemonsInteractor struct {
		repository repository.ISVPokemonRepository
		service    service.ISVPokemonService
		presenter  GetPokemonsPresenter
	}
)

func NewGetSwShPokemonsInteractor(
	repository repository.ISwShPokemonRepository,
	service service.ISwShPokemonService,
	presenter GetPokemonsPresenter,
) GetPokemonsUseCase {
	return &GetSwShPokemonsInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func NewGetSVPokemonsInteractor(
	repository repository.ISVPokemonRepository,
	service service.ISVPokemonService,
	presenter GetPokemonsPresenter,
) GetPokemonsUseCase {
	return &GetSVPokemonsInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSwShPokemonsInteractor) Execute(input GetPokemonsInput) (GetPokemonsOutput, error) {
	pokemons, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Pokemon{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(pokemons)), err
}

func (interactor *GetSVPokemonsInteractor) Execute(input GetPokemonsInput) (GetPokemonsOutput, error) {
	pokemons, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Pokemon{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(pokemons)), err
}

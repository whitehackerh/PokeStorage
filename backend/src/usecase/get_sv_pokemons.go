package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSVPokemonsUseCase interface {
		Execute(input GetSVPokemonsInput) (GetSVPokemonsOutput, error)
	}
	GetSVPokemonsInput     struct{}
	GetSVPokemonsPresenter interface {
		Output([]entity.Pokemon) GetSVPokemonsOutput
	}
	GetSVPokemonsOutput struct {
		Pokemons []api_schema.Pokemon `json:"pokemons"`
	}
	GetSVPokemonsInteractor struct {
		repository repository.ISVPokemonRepository
		service    service.ISVPokemonService
		presenter  GetSVPokemonsPresenter
	}
)

func NewGetSVPokemonsInteractor(
	repository repository.ISVPokemonRepository,
	service service.ISVPokemonService,
	presenter GetSVPokemonsPresenter,
) GetSVPokemonsUseCase {
	return &GetSVPokemonsInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSVPokemonsInteractor) Execute(input GetSVPokemonsInput) (GetSVPokemonsOutput, error) {
	pokemons, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Pokemon{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(pokemons)), err
}

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
	GetSVBredPokemonsUseCase interface {
		Execute(input GetSVBredPokemonsInput) (GetSVBredPokemonsOutput, error)
	}
	GetSwShBredPokemonsInput struct {
		UserId string
	}
	GetSVBredPokemonsInput struct {
		UserId string
	}
	GetSwShBredPokemonsPresenter interface {
		Output([]entity.SwShBredPokemon) GetSwShBredPokemonsOutput
	}
	GetSVBredPokemonsPresenter interface {
		Output([]entity.SVBredPokemon) GetSVBredPokemonsOutput
	}
	GetSwShBredPokemonsOutput struct {
		BredPokemons []api_schema.SwShBredPokemon `json:"bred_pokemons"`
	}
	GetSVBredPokemonsOutput struct {
		BredPokemons []api_schema.SVBredPokemon `json:"bred_pokemons"`
	}
	GetSwShBredPokemonsInteractor struct {
		service    service.ISwShBredPokemonService
		repository repository.ISwShBredPokemonRepository
		presenter  GetSwShBredPokemonsPresenter
	}
	GetSVBredPokemonsInteractor struct {
		service    service.ISVBredPokemonService
		repository repository.ISVBredPokemonRepository
		presenter  GetSVBredPokemonsPresenter
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

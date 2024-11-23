package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PostSwShBredPokemonsUseCase interface {
		Execute(input PostSwShBredPokemonsInput, userId string, tx *gorm.DB) (PostSwShBredPokemonsOutput, error)
	}
	PostSwShBredPokemonsInput struct {
		BredPokemon api_schema.SwShBredPokemon `json:"bred_pokemon"`
	}
	PostSwShBredPokemonsPresenter interface {
		Output() PostSwShBredPokemonsOutput
	}
	PostSwShBredPokemonsOutput     struct{}
	PostSwShBredPokemonsInteractor struct {
		service        service.ISwShBredPokemonService
		bredRepo       repository.ISwShBredPokemonRepository
		individualRepo repository.ISwShIndividualValuesRepository
		baseRepo       repository.ISwShBasePointsRepository
		actualRepo     repository.ISwShActualValuesRepository
		presenter      PostSwShBredPokemonsPresenter
	}
)

func NewPostSwShBredPokemonsInteractor(
	service service.ISwShBredPokemonService,
	bredRepo repository.ISwShBredPokemonRepository,
	individualRepo repository.ISwShIndividualValuesRepository,
	baseRepo repository.ISwShBasePointsRepository,
	actualRepo repository.ISwShActualValuesRepository,
	presenter PostSwShBredPokemonsPresenter,
) PostSwShBredPokemonsUseCase {
	return &PostSwShBredPokemonsInteractor{
		service:        service,
		bredRepo:       bredRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *PostSwShBredPokemonsInteractor) Execute(input PostSwShBredPokemonsInput, userId string, tx *gorm.DB) (PostSwShBredPokemonsOutput, error) {
	bredPokemon := interactor.service.MakeEntityFromApiSchema(input.BredPokemon, userId)

	if err := interactor.bredRepo.Create(
		tx,
		converter.SwShBredPokemonEntityToModel(bredPokemon),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.individualRepo.Create(
		tx,
		converter.IndividualValuesEntityToSwShModel(bredPokemon.IndividualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.baseRepo.Create(
		tx,
		converter.BasePointsEntityToSwShModel(bredPokemon.BasePoints()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.actualRepo.Create(
		tx,
		converter.ActualValuesEntityToSwShModel(bredPokemon.ActualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

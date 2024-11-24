package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PutSwShBredPokemonsUseCase interface {
		Execute(input PutSwShBredPokemonsInput, userId string, tx *gorm.DB) (PutSwShBredPokemonsOutput, error)
	}
	PutSwShBredPokemonsInput struct {
		BredPokemon api_schema.SwShBredPokemon `json:"bred_pokemon"`
	}
	PutSwShBredPokemonsPresenter interface {
		Output() PutSwShBredPokemonsOutput
	}
	PutSwShBredPokemonsOutput     struct{}
	PutSwShBredPokemonsInteractor struct {
		service        service.ISwShBredPokemonService
		bredRepo       repository.ISwShBredPokemonRepository
		individualRepo repository.ISwShIndividualValuesRepository
		baseRepo       repository.ISwShBasePointsRepository
		actualRepo     repository.ISwShActualValuesRepository
		presenter      PutSwShBredPokemonsPresenter
	}
)

func NewPutSwShBredPokemonsInteractor(
	service service.ISwShBredPokemonService,
	bredRepo repository.ISwShBredPokemonRepository,
	individualRepo repository.ISwShIndividualValuesRepository,
	baseRepo repository.ISwShBasePointsRepository,
	actualRepo repository.ISwShActualValuesRepository,
	presenter PutSwShBredPokemonsPresenter,
) PutSwShBredPokemonsUseCase {
	return &PutSwShBredPokemonsInteractor{
		service:        service,
		bredRepo:       bredRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *PutSwShBredPokemonsInteractor) Execute(input PutSwShBredPokemonsInput, userId string, tx *gorm.DB) (PutSwShBredPokemonsOutput, error) {
	bredPokemon := interactor.service.MakeEntityFromApiSchema(input.BredPokemon, userId)

	if err := interactor.bredRepo.Update(
		tx,
		converter.SwShBredPokemonEntityToModel(bredPokemon),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.individualRepo.Update(
		tx,
		converter.IndividualValuesEntityToSwShModel(bredPokemon.IndividualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.baseRepo.Update(
		tx,
		converter.BasePointsEntityToSwShModel(bredPokemon.BasePoints()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.actualRepo.Update(
		tx,
		converter.ActualValuesEntityToSwShModel(bredPokemon.ActualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

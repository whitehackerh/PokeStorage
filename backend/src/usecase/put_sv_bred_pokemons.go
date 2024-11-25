package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PutSVBredPokemonsUseCase interface {
		Execute(input PutSVBredPokemonsInput, userId string, tx *gorm.DB) (PutSVBredPokemonsOutput, error)
	}
	PutSVBredPokemonsInput struct {
		BredPokemon api_schema.SVBredPokemon `json:"bred_pokemon"`
	}
	PutSVBredPokemonsPresenter interface {
		Output() PutSVBredPokemonsOutput
	}
	PutSVBredPokemonsOutput     struct{}
	PutSVBredPokemonsInteractor struct {
		service        service.ISVBredPokemonService
		bredRepo       repository.ISVBredPokemonRepository
		individualRepo repository.ISVIndividualValuesRepository
		baseRepo       repository.ISVBasePointsRepository
		actualRepo     repository.ISVActualValuesRepository
		presenter      PutSVBredPokemonsPresenter
	}
)

func NewPutSVBredPokemonsInteractor(
	service service.ISVBredPokemonService,
	bredRepo repository.ISVBredPokemonRepository,
	individualRepo repository.ISVIndividualValuesRepository,
	baseRepo repository.ISVBasePointsRepository,
	actualRepo repository.ISVActualValuesRepository,
	presenter PutSVBredPokemonsPresenter,
) PutSVBredPokemonsUseCase {
	return &PutSVBredPokemonsInteractor{
		service:        service,
		bredRepo:       bredRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *PutSVBredPokemonsInteractor) Execute(input PutSVBredPokemonsInput, userId string, tx *gorm.DB) (PutSVBredPokemonsOutput, error) {
	bredPokemon := interactor.service.MakeEntityFromApiSchema(input.BredPokemon, userId)

	if err := interactor.bredRepo.Update(
		tx,
		converter.SVBredPokemonEntityToModel(bredPokemon),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.individualRepo.Update(
		tx,
		converter.IndividualValuesEntityToSVModel(bredPokemon.IndividualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.baseRepo.Update(
		tx,
		converter.BasePointsEntityToSVModel(bredPokemon.BasePoints()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.actualRepo.Update(
		tx,
		converter.ActualValuesEntityToSVModel(bredPokemon.ActualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

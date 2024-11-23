package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	PostSVBredPokemonsUseCase interface {
		Execute(input PostSVBredPokemonsInput, userId string, tx *gorm.DB) (PostSVBredPokemonsOutput, error)
	}
	PostSVBredPokemonsInput struct {
		BredPokemon api_schema.SVBredPokemon `json:"bred_pokemon"`
	}
	PostSVBredPokemonsPresenter interface {
		Output() PostSVBredPokemonsOutput
	}
	PostSVBredPokemonsOutput     struct{}
	PostSVBredPokemonsInteractor struct {
		service        service.ISVBredPokemonService
		bredRepo       repository.ISVBredPokemonRepository
		individualRepo repository.ISVIndividualValuesRepository
		baseRepo       repository.ISVBasePointsRepository
		actualRepo     repository.ISVActualValuesRepository
		presenter      PostSVBredPokemonsPresenter
	}
)

func NewPostSVBredPokemonsInteractor(
	service service.ISVBredPokemonService,
	bredRepo repository.ISVBredPokemonRepository,
	individualRepo repository.ISVIndividualValuesRepository,
	baseRepo repository.ISVBasePointsRepository,
	actualRepo repository.ISVActualValuesRepository,
	presenter PostSVBredPokemonsPresenter,
) PostSVBredPokemonsUseCase {
	return &PostSVBredPokemonsInteractor{
		service:        service,
		bredRepo:       bredRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *PostSVBredPokemonsInteractor) Execute(input PostSVBredPokemonsInput, userId string, tx *gorm.DB) (PostSVBredPokemonsOutput, error) {
	bredPokemon := interactor.service.MakeEntityFromApiSchema(input.BredPokemon, userId)

	if err := interactor.bredRepo.Create(
		tx,
		converter.SVBredPokemonEntityToModel(bredPokemon),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.individualRepo.Create(
		tx,
		converter.IndividualValuesEntityToSVModel(bredPokemon.IndividualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.baseRepo.Create(
		tx,
		converter.BasePointsEntityToSVModel(bredPokemon.BasePoints()),
	); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.actualRepo.Create(
		tx,
		converter.ActualValuesEntityToSVModel(bredPokemon.ActualValues()),
	); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

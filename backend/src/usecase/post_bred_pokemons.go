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
		Execute(input PostSwShBredPokemonsInput, userId string, tx *gorm.DB) (PostBredPokemonsOutput, error)
	}
	PostSVBredPokemonsUseCase interface {
		Execute(input PostSVBredPokemonsInput, userId string, tx *gorm.DB) (PostBredPokemonsOutput, error)
	}
	PostSwShBredPokemonsInput struct {
		BredPokemon api_schema.SwShBredPokemon `json:"bred_pokemon"`
	}
	PostSVBredPokemonsInput struct {
		BredPokemon api_schema.SVBredPokemon `json:"bred_pokemon"`
	}
	PostBredPokemonsPresenter interface {
		Output() PostBredPokemonsOutput
	}
	PostBredPokemonsOutput         struct{}
	PostSwShBredPokemonsInteractor struct {
		service        service.ISwShBredPokemonService
		bredRepo       repository.ISwShBredPokemonRepository
		individualRepo repository.ISwShIndividualValuesRepository
		baseRepo       repository.ISwShBasePointsRepository
		actualRepo     repository.ISwShActualValuesRepository
		presenter      PostBredPokemonsPresenter
	}
	PostSVBredPokemonsInteractor struct {
		service        service.ISVBredPokemonService
		bredRepo       repository.ISVBredPokemonRepository
		individualRepo repository.ISVIndividualValuesRepository
		baseRepo       repository.ISVBasePointsRepository
		actualRepo     repository.ISVActualValuesRepository
		presenter      PostBredPokemonsPresenter
	}
)

func NewPostSwShBredPokemonsInteractor(
	service service.ISwShBredPokemonService,
	bredRepo repository.ISwShBredPokemonRepository,
	individualRepo repository.ISwShIndividualValuesRepository,
	baseRepo repository.ISwShBasePointsRepository,
	actualRepo repository.ISwShActualValuesRepository,
	presenter PostBredPokemonsPresenter,
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

func NewPostSVBredPokemonsInteractor(
	service service.ISVBredPokemonService,
	bredRepo repository.ISVBredPokemonRepository,
	individualRepo repository.ISVIndividualValuesRepository,
	baseRepo repository.ISVBasePointsRepository,
	actualRepo repository.ISVActualValuesRepository,
	presenter PostBredPokemonsPresenter,
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

func (interactor *PostSwShBredPokemonsInteractor) Execute(input PostSwShBredPokemonsInput, userId string, tx *gorm.DB) (PostBredPokemonsOutput, error) {
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

func (interactor *PostSVBredPokemonsInteractor) Execute(input PostSVBredPokemonsInput, userId string, tx *gorm.DB) (PostBredPokemonsOutput, error) {
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

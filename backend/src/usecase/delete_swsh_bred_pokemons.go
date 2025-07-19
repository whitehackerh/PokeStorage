package usecase

import (
	"errors"

	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	DeleteSwShBredPokemonsUseCase interface {
		Execute(input DeleteSwShBredPokemonsInput, tx *gorm.DB) (DeleteSwShBredPokemonsOutput, error)
	}
	DeleteSwShBredPokemonsInput struct {
		BredPokemonId string
	}
	DeleteSwShBredPokemonsPresenter interface {
		Output() DeleteSwShBredPokemonsOutput
	}
	DeleteSwShBredPokemonsOutput     struct{}
	DeleteSwShBredPokemonsInteractor struct {
		bredRepo       repository.ISwShBredPokemonRepository
		teamRepo       repository.ISwShTeamRepository
		individualRepo repository.ISwShIndividualValuesRepository
		baseRepo       repository.ISwShBasePointsRepository
		actualRepo     repository.ISwShActualValuesRepository
		presenter      DeleteSwShBredPokemonsPresenter
	}
)

func NewDeleteSwShBredPokemonsInteractor(
	bredRepo repository.ISwShBredPokemonRepository,
	teamRepo repository.ISwShTeamRepository,
	individualRepo repository.ISwShIndividualValuesRepository,
	baseRepo repository.ISwShBasePointsRepository,
	actualRepo repository.ISwShActualValuesRepository,
	presenter DeleteSwShBredPokemonsPresenter,
) DeleteSwShBredPokemonsUseCase {
	return &DeleteSwShBredPokemonsInteractor{
		bredRepo:       bredRepo,
		teamRepo:       teamRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *DeleteSwShBredPokemonsInteractor) Execute(input DeleteSwShBredPokemonsInput, tx *gorm.DB) (DeleteSwShBredPokemonsOutput, error) {
	bredPokemon, err := interactor.bredRepo.FindById(input.BredPokemonId)
	if err != nil {
		return interactor.presenter.Output(), err
	}

	if exists, err := interactor.teamRepo.ExistsBredPokemon(bredPokemon.Id); err != nil {
		return interactor.presenter.Output(), err
	} else if exists {
		return interactor.presenter.Output(), errors.New("this Pokémon is registered with the team and cannot be deleted")
	}

	if err := interactor.individualRepo.Delete(tx, bredPokemon.IndividualValuesId); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.baseRepo.Delete(tx, bredPokemon.BasePointsId); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.actualRepo.Delete(tx, bredPokemon.ActualValuesId); err != nil {
		return interactor.presenter.Output(), err
	}
	if err := interactor.bredRepo.Delete(tx, bredPokemon.Id); err != nil {
		return interactor.presenter.Output(), err
	}
	return interactor.presenter.Output(), nil
}

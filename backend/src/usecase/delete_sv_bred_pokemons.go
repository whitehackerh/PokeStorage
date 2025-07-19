package usecase

import (
	"errors"

	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	DeleteSVBredPokemonsUseCase interface {
		Execute(input DeleteSVBredPokemonsInput, tx *gorm.DB) (DeleteSVBredPokemonsOutput, error)
	}
	DeleteSVBredPokemonsInput struct {
		BredPokemonId string
	}
	DeleteSVBredPokemonsPresenter interface {
		Output() DeleteSVBredPokemonsOutput
	}
	DeleteSVBredPokemonsOutput     struct{}
	DeleteSVBredPokemonsInteractor struct {
		bredRepo       repository.ISVBredPokemonRepository
		teamRepo       repository.ISVTeamRepository
		individualRepo repository.ISVIndividualValuesRepository
		baseRepo       repository.ISVBasePointsRepository
		actualRepo     repository.ISVActualValuesRepository
		presenter      DeleteSVBredPokemonsPresenter
	}
)

func NewDeleteSVBredPokemonsInteractor(
	bredRepo repository.ISVBredPokemonRepository,
	teamRepo repository.ISVTeamRepository,
	individualRepo repository.ISVIndividualValuesRepository,
	baseRepo repository.ISVBasePointsRepository,
	actualRepo repository.ISVActualValuesRepository,
	presenter DeleteSVBredPokemonsPresenter,
) DeleteSVBredPokemonsUseCase {
	return &DeleteSVBredPokemonsInteractor{
		bredRepo:       bredRepo,
		teamRepo:       teamRepo,
		individualRepo: individualRepo,
		baseRepo:       baseRepo,
		actualRepo:     actualRepo,
		presenter:      presenter,
	}
}

func (interactor *DeleteSVBredPokemonsInteractor) Execute(input DeleteSVBredPokemonsInput, tx *gorm.DB) (DeleteSVBredPokemonsOutput, error) {
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

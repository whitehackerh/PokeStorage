package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	DeleteSwShTeamsUseCase interface {
		Execute(input DeleteSwShTeamsInput, tx *gorm.DB) (DeleteSwShTeamsOutput, error)
	}
	DeleteSwShTeamsInput struct {
		TeamId string
	}
	DeleteSwShTeamsPresenter interface {
		Output() DeleteSwShTeamsOutput
	}
	DeleteSwShTeamsOutput     struct{}
	DeleteSwShTeamsInteractor struct {
		repo      repository.ISwShTeamRepository
		presenter DeleteSwShTeamsPresenter
	}
)

func NewDeleteSwShTeamsInteractor(
	repo repository.ISwShTeamRepository,
	presenter DeleteSwShTeamsPresenter,
) DeleteSwShTeamsUseCase {
	return &DeleteSwShTeamsInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (interactor *DeleteSwShTeamsInteractor) Execute(input DeleteSwShTeamsInput, tx *gorm.DB) (DeleteSwShTeamsOutput, error) {
	team, err := interactor.repo.FindById(input.TeamId)
	if err != nil {
		return interactor.presenter.Output(), err
	}

	if err := interactor.repo.Delete(tx, team.Id); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

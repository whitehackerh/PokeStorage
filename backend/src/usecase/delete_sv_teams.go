package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/repository"
	"gorm.io/gorm"
)

type (
	DeleteSVTeamsUseCase interface {
		Execute(input DeleteSVTeamsInput, tx *gorm.DB) (DeleteSVTeamsOutput, error)
	}
	DeleteSVTeamsInput struct {
		TeamId string
	}
	DeleteSVTeamsPresenter interface {
		Output() DeleteSVTeamsOutput
	}
	DeleteSVTeamsOutput     struct{}
	DeleteSVTeamsInteractor struct {
		repo      repository.ISVTeamRepository
		presenter DeleteSVTeamsPresenter
	}
)

func NewDeleteSVTeamsInteractor(
	repo repository.ISVTeamRepository,
	presenter DeleteSVTeamsPresenter,
) DeleteSVTeamsUseCase {
	return &DeleteSVTeamsInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (interactor *DeleteSVTeamsInteractor) Execute(input DeleteSVTeamsInput, tx *gorm.DB) (DeleteSVTeamsOutput, error) {
	team, err := interactor.repo.FindById(input.TeamId)
	if err != nil {
		return interactor.presenter.Output(), err
	}

	if err := interactor.repo.Delete(tx, team.Id); err != nil {
		return interactor.presenter.Output(), err
	}

	return interactor.presenter.Output(), nil
}

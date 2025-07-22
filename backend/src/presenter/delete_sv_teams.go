package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type DeleteSVTeamsPresenter struct{}

func NewDeleteSVTeamsPresenter() usecase.DeleteSVTeamsPresenter {
	return &DeleteSVTeamsPresenter{}
}

func (a *DeleteSVTeamsPresenter) Output() usecase.DeleteSVTeamsOutput {
	return usecase.DeleteSVTeamsOutput{}
}

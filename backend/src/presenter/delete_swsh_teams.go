package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type DeleteSwShTeamsPresenter struct{}

func NewDeleteSwShTeamsPresenter() usecase.DeleteSwShTeamsPresenter {
	return &DeleteSwShTeamsPresenter{}
}

func (a *DeleteSwShTeamsPresenter) Output() usecase.DeleteSwShTeamsOutput {
	return usecase.DeleteSwShTeamsOutput{}
}

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type DeleteSwShBredPokemonsPresenter struct{}

func NewDeleteSwShBredPokemonsPresenter() usecase.DeleteSwShBredPokemonsPresenter {
	return &DeleteSwShBredPokemonsPresenter{}
}

func (a *DeleteSwShBredPokemonsPresenter) Output() usecase.DeleteSwShBredPokemonsOutput {
	return usecase.DeleteSwShBredPokemonsOutput{}
}

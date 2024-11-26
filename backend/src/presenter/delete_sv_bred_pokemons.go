package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type DeleteSVBredPokemonsPresenter struct{}

func NewDeleteSVBredPokemonsPresenter() usecase.DeleteSVBredPokemonsPresenter {
	return &DeleteSVBredPokemonsPresenter{}
}

func (a *DeleteSVBredPokemonsPresenter) Output() usecase.DeleteSVBredPokemonsOutput {
	return usecase.DeleteSVBredPokemonsOutput{}
}

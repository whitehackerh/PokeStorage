package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PutSVBredPokemonsPresenter struct{}

func NewPutSVBredPokemonsPresenter() usecase.PutSVBredPokemonsPresenter {
	return &PutSVBredPokemonsPresenter{}
}

func (a *PutSVBredPokemonsPresenter) Output() usecase.PutSVBredPokemonsOutput {
	return usecase.PutSVBredPokemonsOutput{}
}

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PutSwShBredPokemonsPresenter struct{}

func NewPutSwShBredPokemonsPresenter() usecase.PutSwShBredPokemonsPresenter {
	return &PutSwShBredPokemonsPresenter{}
}

func (a *PutSwShBredPokemonsPresenter) Output() usecase.PutSwShBredPokemonsOutput {
	return usecase.PutSwShBredPokemonsOutput{}
}

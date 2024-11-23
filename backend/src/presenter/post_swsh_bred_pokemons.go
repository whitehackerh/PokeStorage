package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PostSwShBredPokemonsPresenter struct{}

func NewPostSwShBredPokemonsPresenter() usecase.PostSwShBredPokemonsPresenter {
	return &PostSwShBredPokemonsPresenter{}
}

func (a *PostSwShBredPokemonsPresenter) Output() usecase.PostSwShBredPokemonsOutput {
	return usecase.PostSwShBredPokemonsOutput{}
}

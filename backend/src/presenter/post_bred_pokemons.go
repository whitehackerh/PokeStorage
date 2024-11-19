package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PostBredPokemonsPresenter struct{}

func NewPostBredPokemonsPresenter() usecase.PostBredPokemonsPresenter {
	return &PostBredPokemonsPresenter{}
}

func (a *PostBredPokemonsPresenter) Output() usecase.PostBredPokemonsOutput {
	return usecase.PostBredPokemonsOutput{}
}

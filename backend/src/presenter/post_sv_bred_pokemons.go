package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PostSVBredPokemonsPresenter struct{}

func NewPostSVBredPokemonsPresenter() usecase.PostSVBredPokemonsPresenter {
	return &PostSVBredPokemonsPresenter{}
}

func (a *PostSVBredPokemonsPresenter) Output() usecase.PostSVBredPokemonsOutput {
	return usecase.PostSVBredPokemonsOutput{}
}

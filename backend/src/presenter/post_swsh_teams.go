package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PostSwShTeamsPresenter struct{}

func NewPostSwShTeamsPresenter() usecase.PostSwShTeamsPresenter {
	return &PostSwShTeamsPresenter{}
}

func (a *PostSwShTeamsPresenter) Output() usecase.PostSwShTeamsOutput {
	return usecase.PostSwShTeamsOutput{}
}

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PutSwShTeamsPresenter struct{}

func NewPutSwShTeamsPresenter() usecase.PutSwShTeamsPresenter {
	return &PutSwShTeamsPresenter{}
}

func (a *PutSwShTeamsPresenter) Output() usecase.PutSwShTeamsOutput {
	return usecase.PutSwShTeamsOutput{}
}

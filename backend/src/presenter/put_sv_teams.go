package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PutSVTeamsPresenter struct{}

func NewPutSVTeamsPresenter() usecase.PutSVTeamsPresenter {
	return &PutSVTeamsPresenter{}
}

func (a *PutSVTeamsPresenter) Output() usecase.PutSVTeamsOutput {
	return usecase.PutSVTeamsOutput{}
}

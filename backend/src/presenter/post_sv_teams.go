package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type PostSVTeamsPresenter struct{}

func NewPostSVTeamsPresenter() usecase.PostSVTeamsPresenter {
	return &PostSVTeamsPresenter{}
}

func (a *PostSVTeamsPresenter) Output() usecase.PostSVTeamsOutput {
	return usecase.PostSVTeamsOutput{}
}

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type SignInPresenter struct{}

func NewSignInPresenter() usecase.SignInPresenter {
	return &SignInPresenter{}
}

func (a *SignInPresenter) Output() usecase.SignInOutput {
	return usecase.SignInOutput{}
}

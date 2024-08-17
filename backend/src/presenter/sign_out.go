package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type SignOutPresenter struct{}

func NewSignOutPresenter() usecase.SignOutPresenter {
	return &SignOutPresenter{}
}

func (a *SignOutPresenter) Output() usecase.SignOutOutput {
	return usecase.SignOutOutput{}
}

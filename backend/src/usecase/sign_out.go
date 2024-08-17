package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/middleware"
)

type (
	SignOutUseCase interface {
		Execute(input SignOutInput) (SignOutOutput, error)
	}
	SignOutInput struct {
		Token  string
		Claims middleware.Claims
	}
	SignOutPresenter interface {
		Output() SignOutOutput
	}
	SignOutOutput     struct{}
	SignOutInteractor struct {
		auth      middleware.IAuth
		presenter SignOutPresenter
	}
)

func NewSignOutInteractor(
	auth middleware.IAuth,
	presenter SignOutPresenter,
) SignOutUseCase {
	return &SignOutInteractor{
		auth:      auth,
		presenter: presenter,
	}
}

func (interactor *SignOutInteractor) Execute(input SignOutInput) (SignOutOutput, error) {
	err := interactor.auth.CreateBlacklist(input.Token, input.Claims)
	if err != nil {
		return interactor.presenter.Output(), err
	}
	return interactor.presenter.Output(), err
}

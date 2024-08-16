package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/middleware"
)

type (
	SignInUseCase interface {
		Execute(input SignInInput) (SignInOutput, string, error)
	}
	SignInInput struct {
		Username string `json:"username" binding:"required,max=50"`
		Password string `json:"password" binding:"required"`
	}
	SignInPresenter interface {
		Output() SignInOutput
	}
	SignInOutput     struct{}
	SignInInteractor struct {
		service   entity.IUserService
		auth      middleware.IAuth
		presenter SignInPresenter
	}
)

func NewSignInInteractor(
	service entity.IUserService,
	auth middleware.IAuth,
	presenter SignInPresenter,
) SignInUseCase {
	return &SignInInteractor{
		service:   service,
		auth:      auth,
		presenter: presenter,
	}
}

func (interactor *SignInInteractor) Execute(input SignInInput) (SignInOutput, string, error) {
	user, err := interactor.service.FindByUsernameAndPassword(input.Username, input.Password)
	if err != nil {
		return interactor.presenter.Output(), "", err
	}
	token, err := interactor.auth.CreateToken(user)
	if err != nil {
		return interactor.presenter.Output(), "", err
	}
	return interactor.presenter.Output(), token, err
}

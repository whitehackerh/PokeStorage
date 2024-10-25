package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/repository"
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
		repository repository.IUserRepository
		auth       middleware.IAuth
		presenter  SignInPresenter
	}
)

func NewSignInInteractor(
	repository repository.IUserRepository,
	auth middleware.IAuth,
	presenter SignInPresenter,
) SignInUseCase {
	return &SignInInteractor{
		repository: repository,
		auth:       auth,
		presenter:  presenter,
	}
}

func (interactor *SignInInteractor) Execute(input SignInInput) (SignInOutput, string, error) {
	user, err := interactor.repository.FindByUsernameAndPassword(input.Username, input.Password)
	if err != nil {
		return interactor.presenter.Output(), "", err
	}
	token, err := interactor.auth.CreateToken(converter.UserModelToEntity(user))
	if err != nil {
		return interactor.presenter.Output(), "", err
	}
	return interactor.presenter.Output(), token, err
}

package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/util"
)

type (
	SignupUseCase interface {
		Execute(input SignupInput) (SignupOutput, string, error)
	}
	SignupInput struct {
		Username     string `json:"username" binding:"required,max=50"`
		Password     string `json:"password" binding:"required"`
		EmailAddress string `json:"email_address" binding:"required,email"`
		Name         string `json:"name" binding:"required,max=50"`
	}
	SignupPresenter interface {
		Output(entity.User) SignupOutput
	}
	SignupOutput struct {
		Id           string `json:"id"`
		Username     string `json:"username"`
		EmailAddress string `json:"email_address"`
		Name         string `json:"name"`
	}
	SignupInteractor struct {
		service   entity.IUserService
		auth      middleware.IAuth
		presenter SignupPresenter
	}
)

func NewSignupInteractor(
	service entity.IUserService,
	auth middleware.IAuth,
	presenter SignupPresenter,
) SignupUseCase {
	return &SignupInteractor{
		service:   service,
		auth:      auth,
		presenter: presenter,
	}
}

func (interactor *SignupInteractor) Execute(input SignupInput) (SignupOutput, string, error) {
	user := entity.NewUser(
		util.NewUUID(),
		input.Username,
		input.Password,
		input.EmailAddress,
		input.Name,
	)
	user, err := interactor.service.Create(user)
	if err != nil {
		return interactor.presenter.Output(entity.User{}), "", err
	}
	token, err := interactor.auth.CreateToken(user)
	if err != nil {
		return interactor.presenter.Output(entity.User{}), "", err
	}
	return interactor.presenter.Output(user), token, err
}

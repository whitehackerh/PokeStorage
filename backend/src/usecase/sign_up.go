package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/util"
)

type (
	SignUpUseCase interface {
		Execute(input SignUpInput) (SignUpOutput, string, error)
	}
	SignUpInput struct {
		Username     string `json:"username" binding:"required,max=50"`
		Password     string `json:"password" binding:"required"`
		EmailAddress string `json:"email_address" binding:"required,email"`
		Name         string `json:"name" binding:"required,max=50"`
	}
	SignUpPresenter interface {
		Output(entity.User) SignUpOutput
	}
	SignUpOutput struct {
		// TODO Delete field after develop other api
		Id           string `json:"id"`
		Username     string `json:"username"`
		EmailAddress string `json:"email_address"`
		Name         string `json:"name"`
	}
	SignUpInteractor struct {
		service   service.IUserService
		auth      middleware.IAuth
		presenter SignUpPresenter
	}
)

func NewSignUpInteractor(
	service service.IUserService,
	auth middleware.IAuth,
	presenter SignUpPresenter,
) SignUpUseCase {
	return &SignUpInteractor{
		service:   service,
		auth:      auth,
		presenter: presenter,
	}
}

func (interactor *SignUpInteractor) Execute(input SignUpInput) (SignUpOutput, string, error) {
	user := entity.NewUser(
		util.NewUUID(),
		input.Username,
		input.Password,
		input.EmailAddress,
		input.Name,
	)
	err := interactor.service.Create(user)
	if err != nil {
		return interactor.presenter.Output(entity.User{}), "", err
	}
	token, err := interactor.auth.CreateToken(user)
	if err != nil {
		return interactor.presenter.Output(entity.User{}), "", err
	}
	return interactor.presenter.Output(user), token, err
}

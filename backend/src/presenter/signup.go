package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type signupPresenter struct{}

func NewSignupPresenter() usecase.SignupPresenter {
	return signupPresenter{}
}

func (a signupPresenter) Output(user entity.User) usecase.SignupOutput {
	return usecase.SignupOutput{
		Id:           user.Id(),
		Username:     user.Username(),
		EmailAddress: user.EmailAddress(),
		Name:         user.Name(),
	}
}

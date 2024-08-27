package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type SignUpPresenter struct{}

func NewSignUpPresenter() usecase.SignUpPresenter {
	return &SignUpPresenter{}
}

func (p *SignUpPresenter) Output(user entity.User) usecase.SignUpOutput {
	return usecase.SignUpOutput{
		Id:           user.Id(),
		Username:     user.Username(),
		EmailAddress: user.EmailAddress(),
		Name:         user.Name(),
	}
}

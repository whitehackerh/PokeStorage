package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetGendersPresenter struct{}

func NewGetGendersPresenter() usecase.GetGendersPresenter {
	return &GetGendersPresenter{}
}

func (p *GetGendersPresenter) Output(genders []entity.Gender) usecase.GetGendersOutput {
	output := usecase.GetGendersOutput{
		Genders: []api_schema.Gender{},
	}
	for _, gender := range genders {
		output.Genders = append(output.Genders, api_schema.Gender{
			Id:   gender.Id(),
			Name: gender.Name(),
		})
	}
	return output
}

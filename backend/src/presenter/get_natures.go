package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetNaturesPresenter struct{}

func NewGetNaturesPresenter() usecase.GetNaturesPresenter {
	return &GetNaturesPresenter{}
}

func (p *GetNaturesPresenter) Output(natures []entity.Nature) usecase.GetNaturesOutput {
	output := usecase.GetNaturesOutput{
		Natures: []api_schema.Nature{},
	}
	for _, nature := range natures {
		output.Natures = append(output.Natures, api_schema.Nature{
			Id:   nature.Id(),
			Name: nature.Name(),
		})
	}
	return output
}

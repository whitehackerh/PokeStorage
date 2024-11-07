package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetTeraTypesPresenter struct{}

func NewGetTeraTypesPresenter() usecase.GetTeraTypesPresenter {
	return &GetTeraTypesPresenter{}
}

func (p *GetTeraTypesPresenter) Output(teraTypes []entity.TeraType) usecase.GetTeraTypesOutput {
	output := usecase.GetTeraTypesOutput{
		TeraTypes: []api_schema.TeraType{},
	}
	for _, teraType := range teraTypes {
		output.TeraTypes = append(output.TeraTypes, api_schema.TeraType{
			Id:   teraType.Id(),
			Name: teraType.Name(),
		})
	}
	return output
}

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/response_component"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetTeraTypesPresenter struct{}

func NewGetTeraTypesPresenter() usecase.GetTeraTypesPresenter {
	return &GetTeraTypesPresenter{}
}

func (p *GetTeraTypesPresenter) Output(teraTypes []entity.TeraType) usecase.GetTeraTypesOutput {
	output := usecase.GetTeraTypesOutput{
		TeraTypes: []response_component.TeraType{},
	}
	for _, teraType := range teraTypes {
		output.TeraTypes = append(output.TeraTypes, response_component.TeraType{
			Id:   teraType.Id(),
			Name: teraType.Name(),
		})
	}
	return output
}

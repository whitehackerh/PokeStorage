package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetSwShItemsPresenter struct{}

func NewGetSwShItemsPresenter() usecase.GetSwShItemsPresenter {
	return &GetSwShItemsPresenter{}
}

func (p *GetSwShItemsPresenter) Output(Items []entity.Item) usecase.GetSwShItemsOutput {
	output := usecase.GetSwShItemsOutput{
		Items: []api_schema.Item{},
	}
	for _, item := range Items {
		output.Items = append(output.Items, api_schema.Item{
			Id:   item.Id(),
			Name: item.Name(),
		})
	}
	return output
}

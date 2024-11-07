package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetItemsPresenter struct{}

func NewGetItemsPresenter() usecase.GetItemsPresenter {
	return &GetItemsPresenter{}
}

func (p *GetItemsPresenter) Output(Items []entity.Item) usecase.GetItemsOutput {
	output := usecase.GetItemsOutput{
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

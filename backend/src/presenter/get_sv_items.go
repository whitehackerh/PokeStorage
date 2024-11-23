package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetSVItemsPresenter struct{}

func NewGetSVItemsPresenter() usecase.GetSVItemsPresenter {
	return &GetSVItemsPresenter{}
}

func (p *GetSVItemsPresenter) Output(Items []entity.Item) usecase.GetSVItemsOutput {
	output := usecase.GetSVItemsOutput{
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

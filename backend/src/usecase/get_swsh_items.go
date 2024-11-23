package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSwShItemsUseCase interface {
		Execute(input GetSwShItemsInput) (GetSwShItemsOutput, error)
	}
	GetSwShItemsInput     struct{}
	GetSwShItemsPresenter interface {
		Output([]entity.Item) GetSwShItemsOutput
	}
	GetSwShItemsOutput struct {
		Items []api_schema.Item `json:"items"`
	}
	GetSwShItemsInteractor struct {
		repository repository.ISwShItemRepository
		presenter  GetSwShItemsPresenter
	}
)

func NewSwShGetItemsInteractor(
	repository repository.ISwShItemRepository,
	presenter GetSwShItemsPresenter,
) GetSwShItemsUseCase {
	return &GetSwShItemsInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetSwShItemsInteractor) Execute(input GetSwShItemsInput) (GetSwShItemsOutput, error) {
	items, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Item{}), err
	}
	return interactor.presenter.Output(converter.SwShItemModelsToEntities(items)), err
}

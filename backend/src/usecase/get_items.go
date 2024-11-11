package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetItemsUseCase interface {
		Execute(input GetItemsInput) (GetItemsOutput, error)
	}
	GetItemsInput     struct{}
	GetItemsPresenter interface {
		Output([]entity.Item) GetItemsOutput
	}
	GetItemsOutput struct {
		Items []api_schema.Item `json:"items"`
	}
	GetSwShItemsInteractor struct {
		repository repository.ISwShItemRepository
		presenter  GetItemsPresenter
	}
	GetSVItemsInteractor struct {
		repository repository.ISVItemRepository
		presenter  GetItemsPresenter
	}
)

func NewSwShGetItemsInteractor(
	repository repository.ISwShItemRepository,
	presenter GetItemsPresenter,
) GetItemsUseCase {
	return &GetSwShItemsInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func NewSVGetItemsInteractor(
	repository repository.ISVItemRepository,
	presenter GetItemsPresenter,
) GetItemsUseCase {
	return &GetSVItemsInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetSwShItemsInteractor) Execute(input GetItemsInput) (GetItemsOutput, error) {
	items, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Item{}), err
	}
	return interactor.presenter.Output(converter.SwShItemModelsToEntities(items)), err
}

func (interactor *GetSVItemsInteractor) Execute(input GetItemsInput) (GetItemsOutput, error) {
	items, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Item{}), err
	}
	return interactor.presenter.Output(converter.SVItemModelsToEntities(items)), err
}

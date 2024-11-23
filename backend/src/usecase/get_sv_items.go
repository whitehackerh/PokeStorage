package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSVItemsUseCase interface {
		Execute(input GetSVItemsInput) (GetSVItemsOutput, error)
	}
	GetSVItemsInput     struct{}
	GetSVItemsPresenter interface {
		Output([]entity.Item) GetSVItemsOutput
	}
	GetSVItemsOutput struct {
		Items []api_schema.Item `json:"items"`
	}
	GetSVItemsInteractor struct {
		repository repository.ISVItemRepository
		presenter  GetSVItemsPresenter
	}
)

func NewSVGetItemsInteractor(
	repository repository.ISVItemRepository,
	presenter GetSVItemsPresenter,
) GetSVItemsUseCase {
	return &GetSVItemsInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetSVItemsInteractor) Execute(input GetSVItemsInput) (GetSVItemsOutput, error) {
	items, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Item{}), err
	}
	return interactor.presenter.Output(converter.SVItemModelsToEntities(items)), err
}

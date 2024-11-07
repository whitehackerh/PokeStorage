package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/response_component"
)

type (
	GetTeraTypesUseCase interface {
		Execute(input GetTeraTypesInput) (GetTeraTypesOutput, error)
	}
	GetTeraTypesInput     struct{}
	GetTeraTypesPresenter interface {
		Output([]entity.TeraType) GetTeraTypesOutput
	}
	GetTeraTypesOutput struct {
		TeraTypes []response_component.TeraType `json:"tera_types"`
	}
	GetTeraTypesInteractor struct {
		repository repository.ITeraTypeRepository
		presenter  GetTeraTypesPresenter
	}
)

func NewGetTeraTypesInteractor(
	repository repository.ITeraTypeRepository,
	presenter GetTeraTypesPresenter,
) GetTeraTypesUseCase {
	return &GetTeraTypesInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetTeraTypesInteractor) Execute(input GetTeraTypesInput) (GetTeraTypesOutput, error) {
	teraTypes, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.TeraType{}), err
	}
	return interactor.presenter.Output(converter.TeraTypeModelsToEntities(teraTypes)), err
}

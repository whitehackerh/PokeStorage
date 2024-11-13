package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetGendersUseCase interface {
		Execute(input GetGendersInput) (GetGendersOutput, error)
	}
	GetGendersInput     struct{}
	GetGendersPresenter interface {
		Output([]entity.Gender) GetGendersOutput
	}
	GetGendersOutput struct {
		Genders []api_schema.Gender `json:"genders"`
	}
	GetGendersInteractor struct {
		repository repository.IGenderRepository
		presenter  GetGendersPresenter
	}
)

func NewGetGendersInteractor(
	repository repository.IGenderRepository,
	presenter GetGendersPresenter,
) GetGendersUseCase {
	return &GetGendersInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetGendersInteractor) Execute(input GetGendersInput) (GetGendersOutput, error) {
	genders, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Gender{}), err
	}
	return interactor.presenter.Output(converter.GenderModelsToEntities(genders)), err
}

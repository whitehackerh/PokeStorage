package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetNaturesUseCase interface {
		Execute(input GetNaturesInput) (GetNaturesOutput, error)
	}
	GetNaturesInput     struct{}
	GetNaturesPresenter interface {
		Output([]entity.Nature) GetNaturesOutput
	}
	GetNaturesOutput struct {
		Natures []api_schema.Nature `json:"natures"`
	}
	GetNaturesInteractor struct {
		repository repository.INatureRepository
		presenter  GetNaturesPresenter
	}
)

func NewGetNaturesInteractor(
	repository repository.INatureRepository,
	presenter GetNaturesPresenter,
) GetNaturesUseCase {
	return &GetNaturesInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetNaturesInteractor) Execute(input GetNaturesInput) (GetNaturesOutput, error) {
	natures, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Nature{}), err
	}
	return interactor.presenter.Output(converter.NatureModelsToEntities(natures)), err
}

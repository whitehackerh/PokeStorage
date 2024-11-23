package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSVMovesUseCase interface {
		Execute(input GetSVMovesInput) (GetSVMovesOutput, error)
	}
	GetSVMovesInput     struct{}
	GetSVMovesPresenter interface {
		Output([]entity.Move) GetSVMovesOutput
	}
	GetSVMovesOutput struct {
		Moves []api_schema.Move `json:"moves"`
	}
	GetSVMovesInteractor struct {
		repository repository.ISVMoveRepository
		service    service.ISVMoveService
		presenter  GetSVMovesPresenter
	}
)

func NewSVGetMovesInteractor(
	repository repository.ISVMoveRepository,
	service service.ISVMoveService,
	presenter GetSVMovesPresenter,
) GetSVMovesUseCase {
	return &GetSVMovesInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSVMovesInteractor) Execute(input GetSVMovesInput) (GetSVMovesOutput, error) {
	moves, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Move{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(moves)), err
}

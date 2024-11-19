package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetMovesUseCase interface {
		Execute(input GetMovesInput) (GetMovesOutput, error)
	}
	GetMovesInput     struct{}
	GetMovesPresenter interface {
		Output([]entity.Move) GetMovesOutput
	}
	GetMovesOutput struct {
		Moves []api_schema.Move `json:"moves"`
	}
	GetSwShMovesInteractor struct {
		repository repository.ISwShMoveRepository
		service    service.ISwShMoveService
		presenter  GetMovesPresenter
	}
	GetSVMovesInteractor struct {
		repository repository.ISVMoveRepository
		service    service.ISVMoveService
		presenter  GetMovesPresenter
	}
)

func NewSwShGetMovesInteractor(
	repository repository.ISwShMoveRepository,
	service service.ISwShMoveService,
	presenter GetMovesPresenter,
) GetMovesUseCase {
	return &GetSwShMovesInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func NewSVGetMovesInteractor(
	repository repository.ISVMoveRepository,
	service service.ISVMoveService,
	presenter GetMovesPresenter,
) GetMovesUseCase {
	return &GetSVMovesInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSwShMovesInteractor) Execute(input GetMovesInput) (GetMovesOutput, error) {
	moves, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Move{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(moves)), err
}

func (interactor *GetSVMovesInteractor) Execute(input GetMovesInput) (GetMovesOutput, error) {
	moves, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Move{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(moves)), err
}

package usecase

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	GetSwShMovesUseCase interface {
		Execute(input GetSwShMovesInput) (GetSwShMovesOutput, error)
	}
	GetSwShMovesInput     struct{}
	GetSwShMovesPresenter interface {
		Output([]entity.Move) GetSwShMovesOutput
	}
	GetSwShMovesOutput struct {
		Moves []api_schema.Move `json:"moves"`
	}
	GetSwShMovesInteractor struct {
		repository repository.ISwShMoveRepository
		service    service.ISwShMoveService
		presenter  GetSwShMovesPresenter
	}
)

func NewSwShGetMovesInteractor(
	repository repository.ISwShMoveRepository,
	service service.ISwShMoveService,
	presenter GetSwShMovesPresenter,
) GetSwShMovesUseCase {
	return &GetSwShMovesInteractor{
		repository: repository,
		service:    service,
		presenter:  presenter,
	}
}

func (interactor *GetSwShMovesInteractor) Execute(input GetSwShMovesInput) (GetSwShMovesOutput, error) {
	moves, err := interactor.repository.Fetch()
	if err != nil {
		return interactor.presenter.Output([]entity.Move{}), err
	}
	return interactor.presenter.Output(interactor.service.MakeEntitiesFromModels(moves)), err
}

package usecase

import (
	"time"

	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
)

type (
	GetTitlesUseCase interface {
		Execute(input GetTitlesInput) (GetTitlesOutput, error)
	}
	GetTitlesInput     struct{}
	GetTitlesPresenter interface {
		Output([]entity.Title) GetTitlesOutput
	}
	GetTitlesOutput struct {
		Titles []struct {
			Id          int       `json:"id"`
			Title       string    `json:"title"`
			ReleaseDate time.Time `json:"release_date"`
		} `json:"titles"`
	}
	GetTitlesInteractor struct {
		service   service.ITitleService
		presenter GetTitlesPresenter
	}
)

func NewGetTitlesInteractor(
	service service.ITitleService,
	presenter GetTitlesPresenter,
) GetTitlesUseCase {
	return &GetTitlesInteractor{
		service:   service,
		presenter: presenter,
	}
}

func (interactor *GetTitlesInteractor) Execute(input GetTitlesInput) (GetTitlesOutput, error) {
	titles, err := interactor.service.Get()
	if err != nil {
		return interactor.presenter.Output([]entity.Title{}), err
	}
	return interactor.presenter.Output(titles), err
}

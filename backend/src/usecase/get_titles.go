package usecase

import (
	"time"

	"github.com/whitehackerh/PokeStorage/src/converter"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/repository"
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
		repository repository.ITitleRepository
		presenter  GetTitlesPresenter
	}
)

func NewGetTitlesInteractor(
	repository repository.ITitleRepository,
	presenter GetTitlesPresenter,
) GetTitlesUseCase {
	return &GetTitlesInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

func (interactor *GetTitlesInteractor) Execute(input GetTitlesInput) (GetTitlesOutput, error) {
	titles, err := interactor.repository.Get()
	if err != nil {
		return interactor.presenter.Output([]entity.Title{}), err
	}
	return interactor.presenter.Output(converter.TitleModelsToEntities(titles)), err
}

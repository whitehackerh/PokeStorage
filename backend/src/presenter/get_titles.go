package presenter

import (
	"time"

	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetTitlesPresenter struct{}

func NewGetTitlesPresenter() usecase.GetTitlesPresenter {
	return &GetTitlesPresenter{}
}

func (p *GetTitlesPresenter) Output(titles []entity.Title) usecase.GetTitlesOutput {
	output := usecase.GetTitlesOutput{
		Titles: make([]struct {
			Id          int       `json:"id"`
			Title       string    `json:"title"`
			ReleaseDate time.Time `json:"release_date"`
		}, len(titles)),
	}
	for i, title := range titles {
		output.Titles[i] = struct {
			Id          int       `json:"id"`
			Title       string    `json:"title"`
			ReleaseDate time.Time `json:"release_date"`
		}{
			Id:          title.Id(),
			Title:       title.Title(),
			ReleaseDate: title.ReleaseDate(),
		}
	}
	return output
}

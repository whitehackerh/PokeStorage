package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	TitleService struct {
		repo ITitleRepository
	}
	ITitleRepository interface {
		Get() ([]model.Title, error)
	}
)

func NewTitleService(titleRepository ITitleRepository) entity.ITitleService {
	return &TitleService{
		repo: titleRepository,
	}
}

func (t *TitleService) Get() ([]entity.Title, error) {
	titles, err := t.repo.Get()
	if err != nil {
		return []entity.Title{}, err
	}
	return t.mapModelsToEntities(titles), err
}

func (t *TitleService) mapModelsToEntities(titles []model.Title) []entity.Title {
	var entities []entity.Title
	for _, t := range titles {
		entities = append(
			entities,
			entity.NewTitle(t.Id, t.Title, t.ReleaseDate),
		)
	}
	return entities
}

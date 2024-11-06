package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func TitleModelsToEntities(titles []model.Title) []entity.Title {
	var entities []entity.Title
	for _, t := range titles {
		entities = append(entities, TitleModelToEntity(t))
	}
	return entities
}

func TitleModelToEntity(title model.Title) entity.Title {
	return entity.NewTitle(title.Id, title.Title, title.ReleaseDate)
}

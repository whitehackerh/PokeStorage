package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func NatureModelsToEntities(natures []model.Nature) []entity.Nature {
	var entities []entity.Nature
	for _, nature := range natures {
		entities = append(entities, NatureModelToEntity(nature))
	}
	return entities
}

func NatureModelToEntity(nature model.Nature) entity.Nature {
	return entity.NewNature(nature.Id, nature.Name)
}

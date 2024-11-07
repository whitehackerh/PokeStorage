package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func SVItemModelsToEntities(items []model.SVItem) []entity.Item {
	var entities []entity.Item
	for _, t := range items {
		entities = append(entities, SVItemModelToEntity(t))
	}
	return entities
}

func SVItemModelToEntity(item model.SVItem) entity.Item {
	return entity.NewItem(item.Id, item.Name)
}

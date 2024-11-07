package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func SwShItemModelsToEntities(items []model.SwShItem) []entity.Item {
	var entities []entity.Item
	for _, t := range items {
		entities = append(entities, SwShItemModelToEntity(t))
	}
	return entities
}

func SwShItemModelToEntity(item model.SwShItem) entity.Item {
	return entity.NewItem(item.Id, item.Name)
}

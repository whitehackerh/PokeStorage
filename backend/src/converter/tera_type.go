package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func TeraTypeModelsToEntities(teraTypes []model.TeraType) []entity.TeraType {
	var entities []entity.TeraType
	for _, t := range teraTypes {
		entities = append(entities, TeraTypeModelToEntity(t))
	}
	return entities
}

func TeraTypeModelToEntity(teraType model.TeraType) entity.TeraType {
	return entity.NewTeraType(teraType.Id, teraType.Name)
}

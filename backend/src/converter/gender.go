package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func GenderModelsToEntities(genders []model.Gender) []entity.Gender {
	var entities []entity.Gender
	for _, gender := range genders {
		entities = append(entities, GenderModelToEntity(gender))
	}
	return entities
}

func GenderModelToEntity(gender model.Gender) entity.Gender {
	return entity.NewGender(gender.Id, gender.Name)
}

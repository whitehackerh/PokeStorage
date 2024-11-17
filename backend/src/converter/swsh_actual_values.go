package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func ActualValuesEntityToSwShModel(actualValues entity.ActualValues) model.SwShActualValues {
	return model.SwShActualValues{
		Id:             actualValues.Id(),
		BredPokemonId:  actualValues.BredPokemonId(),
		HitPoints:      actualValues.HitPoints(),
		Attack:         actualValues.Attack(),
		Defense:        actualValues.Defense(),
		SpecialAttack:  actualValues.SpecialAttack(),
		SpecialDefense: actualValues.SpecialDefense(),
		Speed:          actualValues.Speed(),
	}
}

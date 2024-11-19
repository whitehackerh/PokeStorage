package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func ActualValuesEntityToSVModel(actualValues entity.ActualValues) model.SVActualValues {
	return model.SVActualValues{
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

package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func IndividualValuesEntityToSwShModel(individualValues entity.IndividualValues) model.SwShIndividualValues {
	return model.SwShIndividualValues{
		Id:             individualValues.Id(),
		BredPokemonId:  individualValues.BredPokemonId(),
		HitPoints:      individualValues.HitPoints(),
		Attack:         individualValues.Attack(),
		Defense:        individualValues.Defense(),
		SpecialAttack:  individualValues.SpecialAttack(),
		SpecialDefense: individualValues.SpecialDefense(),
		Speed:          individualValues.Speed(),
	}
}

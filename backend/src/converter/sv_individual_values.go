package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func IndividualValuesEntityToSVModel(individualValues entity.IndividualValues) model.SVIndividualValues {
	return model.SVIndividualValues{
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

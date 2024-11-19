package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func BasePointsEntityToSwShModel(basePoints entity.BasePoints) model.SwShBasePoints {
	return model.SwShBasePoints{
		Id:             basePoints.Id(),
		BredPokemonId:  basePoints.BredPokemonId(),
		HitPoints:      basePoints.HitPoints(),
		Attack:         basePoints.Attack(),
		Defense:        basePoints.Defense(),
		SpecialAttack:  basePoints.SpecialAttack(),
		SpecialDefense: basePoints.SpecialDefense(),
		Speed:          basePoints.Speed(),
	}
}

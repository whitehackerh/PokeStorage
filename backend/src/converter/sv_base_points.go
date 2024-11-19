package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func BasePointsEntityToSVModel(basePoints entity.BasePoints) model.SVBasePoints {
	return model.SVBasePoints{
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

package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
	"github.com/whitehackerh/PokeStorage/src/util"
)

func SwShBredPokemonEntityToModel(bredPokemon entity.SwShBredPokemon) model.SwShBredPokemon {
	gender := bredPokemon.Gender()
	ability := bredPokemon.Ability()
	nature := bredPokemon.Nature()
	var heldItem *entity.Item = bredPokemon.HeldItem()
	var itemId *int
	if heldItem != nil {
		itemId = util.IntToPointer(heldItem.Id())
	}
	individualValues := bredPokemon.IndividualValues()
	basePoints := bredPokemon.BasePoints()
	actualValues := bredPokemon.ActualValues()
	moveIds := []*int{nil, nil, nil, nil}
	for i, move := range bredPokemon.Moves() {
		moveIds[i] = util.IntToPointer(move.Id())
	}
	return model.SwShBredPokemon{
		Id:                 bredPokemon.Id(),
		UserId:             bredPokemon.UserId(),
		PokemonId:          bredPokemon.PokemonId(),
		GenderId:           gender.Id(),
		Level:              bredPokemon.Level(),
		IsGigantamax:       bredPokemon.IsGigantamax(),
		AbilityId:          ability.Id(),
		NatureId:           nature.Id(),
		ItemId:             itemId,
		IndividualValuesId: individualValues.Id(),
		BasePointsId:       basePoints.Id(),
		ActualValuesId:     actualValues.Id(),
		Move1Id:            bredPokemon.Moves()[0].Id(),
		Move2Id:            moveIds[1],
		Move3Id:            moveIds[2],
		Move4Id:            moveIds[3],
		Note:               bredPokemon.Note(),
	}
}

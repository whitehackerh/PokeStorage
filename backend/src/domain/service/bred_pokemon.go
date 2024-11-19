package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/util"
)

type (
	IBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.BredPokemon, string) entity.BredPokemon
	}
	BredPokemonService struct{}
)

func NewBredPokemonService() IBredPokemonService {
	return &BredPokemonService{}
}

func (b *BredPokemonService) MakeEntityFromApiSchema(schema api_schema.BredPokemon, userId string) entity.BredPokemon {
	var id string
	if schema.Id == "" {
		id = util.NewUUID()
	} else {
		id = schema.Id
	}

	var types []entity.Type
	for _, t := range schema.Types {
		types = append(types, entity.NewType(t.Id, t.Name))
	}

	var individualValuesId string
	if schema.IndividualValues.Id == "" {
		individualValuesId = util.NewUUID()
	} else {
		individualValuesId = schema.IndividualValues.Id
	}

	var basePointsId string
	if schema.BasePoints.Id == "" {
		basePointsId = util.NewUUID()
	} else {
		basePointsId = schema.BasePoints.Id
	}

	var actualValuesId string
	if schema.ActualValues.Id == "" {
		actualValuesId = util.NewUUID()
	} else {
		actualValuesId = schema.ActualValues.Id
	}

	var heldItem *entity.Item
	if schema.HeldItem == nil {
		heldItem = nil
	} else {
		item := entity.NewItem(
			schema.HeldItem.Id,
			schema.HeldItem.Name,
		)
		heldItem = &item
	}

	var moves []entity.Move
	for _, move := range schema.Moves {
		moves = append(moves, entity.NewMove(
			move.Id,
			move.Name,
			entity.NewType(
				move.MoveType.Id,
				move.MoveType.Name,
			),
			entity.NewMoveCategory(
				move.MoveCategory.Id,
				move.MoveCategory.Name,
			),
			move.Power,
			move.Accuracy,
		))
	}

	return entity.NewBredPokemon(
		id,
		userId,
		schema.PokemonId,
		schema.NationalPokedexNo,
		schema.FormeNo,
		schema.Name,
		schema.FormeName,
		entity.NewGender(schema.Gender.Id, schema.Gender.Name),
		50, // schema.Level,
		types,
		entity.NewAbility(schema.Ability.Id, schema.Ability.Name),
		entity.NewNature(schema.Nature.Id, schema.Nature.Name),
		heldItem,
		entity.NewBaseStats(
			schema.BaseStats.Id,
			schema.PokemonId,
			schema.BaseStats.HitPoints,
			schema.BaseStats.Attack,
			schema.BaseStats.Defense,
			schema.BaseStats.SpecialAttack,
			schema.BaseStats.SpecialDefense,
			schema.BaseStats.Speed,
			schema.BaseStats.Average,
			schema.BaseStats.Total,
		),
		entity.NewIndividualValues(
			individualValuesId,
			id,
			schema.IndividualValues.HitPoints,
			schema.IndividualValues.Attack,
			schema.IndividualValues.Defense,
			schema.IndividualValues.SpecialAttack,
			schema.IndividualValues.SpecialDefense,
			schema.IndividualValues.Speed,
		),
		entity.NewBasePoints(
			basePointsId,
			id,
			schema.BasePoints.HitPoints,
			schema.BasePoints.Attack,
			schema.BasePoints.Defense,
			schema.BasePoints.SpecialAttack,
			schema.BasePoints.SpecialDefense,
			schema.BasePoints.Speed,
		),
		entity.NewActualValues(
			actualValuesId,
			id,
			schema.ActualValues.HitPoints,
			schema.ActualValues.Attack,
			schema.ActualValues.Defense,
			schema.ActualValues.SpecialAttack,
			schema.ActualValues.SpecialDefense,
			schema.ActualValues.Speed,
		),
		moves,
		schema.Note,
	)
}

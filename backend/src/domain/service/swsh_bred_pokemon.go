package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.SwShBredPokemon, string) entity.SwShBredPokemon
		MakeEntityFromModel(model.SwShBredPokemonRelation) entity.SwShBredPokemon
	}
	SwShBredPokemonService struct {
		bredPokemonService IBredPokemonService
	}
)

func NewSwShBredPokemonService(bredPokemonService IBredPokemonService) ISwShBredPokemonService {
	return &SwShBredPokemonService{
		bredPokemonService: bredPokemonService,
	}
}

func (s *SwShBredPokemonService) MakeEntityFromApiSchema(schema api_schema.SwShBredPokemon, userId string) entity.SwShBredPokemon {
	return entity.NewSwShBredPokemon(
		s.bredPokemonService.MakeEntityFromApiSchema(schema.BredPokemon, userId),
		schema.IsGigantamax,
	)
}

func (s *SwShBredPokemonService) MakeEntityFromModel(model model.SwShBredPokemonRelation) entity.SwShBredPokemon {
	var types []entity.Type
	types = append(types, entity.NewType(
		model.SwShPokemonRelation.Type1.Id,
		model.SwShPokemonRelation.Type1.Name,
	))
	if model.SwShPokemonRelation.Type2 != nil {
		types = append(types, entity.NewType(
			model.SwShPokemonRelation.Type2.Id,
			model.SwShPokemonRelation.Type2.Name,
		))
	}

	var heldItem *entity.Item
	if model.HeldItem == nil {
		heldItem = nil
	} else {
		item := entity.NewItem(
			model.HeldItem.Id,
			model.HeldItem.Name,
		)
		heldItem = &item
	}

	var moves []entity.Move
	moves = append(moves, entity.NewMove(
		model.Move1Relation.Id,
		model.Move1Relation.Name,
		entity.NewType(
			model.Move1Relation.Type.Id,
			model.Move1Relation.Type.Name,
		),
		entity.NewMoveCategory(
			model.Move1Relation.MoveCategory.Id,
			model.Move1Relation.MoveCategory.Name,
		),
		model.Move1Relation.Power,
		model.Move1Relation.Accuracy,
	))
	if model.Move2Relation != nil {
		moves = append(moves, entity.NewMove(
			model.Move2Relation.Id,
			model.Move2Relation.Name,
			entity.NewType(
				model.Move2Relation.Type.Id,
				model.Move2Relation.Type.Name,
			),
			entity.NewMoveCategory(
				model.Move2Relation.MoveCategory.Id,
				model.Move2Relation.MoveCategory.Name,
			),
			model.Move2Relation.Power,
			model.Move2Relation.Accuracy,
		))
	}
	if model.Move3Relation != nil {
		moves = append(moves, entity.NewMove(
			model.Move3Relation.Id,
			model.Move3Relation.Name,
			entity.NewType(
				model.Move3Relation.Type.Id,
				model.Move3Relation.Type.Name,
			),
			entity.NewMoveCategory(
				model.Move3Relation.MoveCategory.Id,
				model.Move3Relation.MoveCategory.Name,
			),
			model.Move3Relation.Power,
			model.Move3Relation.Accuracy,
		))
	}
	if model.Move4Relation != nil {
		moves = append(moves, entity.NewMove(
			model.Move4Relation.Id,
			model.Move4Relation.Name,
			entity.NewType(
				model.Move4Relation.Type.Id,
				model.Move4Relation.Type.Name,
			),
			entity.NewMoveCategory(
				model.Move4Relation.MoveCategory.Id,
				model.Move4Relation.MoveCategory.Name,
			),
			model.Move4Relation.Power,
			model.Move4Relation.Accuracy,
		))
	}

	return entity.NewSwShBredPokemon(
		entity.NewBredPokemon(
			model.Id,
			model.UserId,
			model.PokemonId,
			model.SwShPokemonRelation.NationalPokedexNo,
			model.SwShPokemonRelation.FormeNo,
			model.SwShPokemonRelation.Name,
			model.SwShPokemonRelation.FormeName,
			entity.NewGender(model.Gender.Id, model.Gender.Name),
			model.Level,
			types,
			entity.NewAbility(model.Ability.Id, model.Ability.Name),
			entity.NewNature(model.Nature.Id, model.Nature.Name),
			heldItem,
			entity.NewBaseStats(
				model.SwShPokemonRelation.BaseStats.Id,
				model.PokemonId,
				model.SwShPokemonRelation.BaseStats.HitPoints,
				model.SwShPokemonRelation.BaseStats.Attack,
				model.SwShPokemonRelation.BaseStats.Defense,
				model.SwShPokemonRelation.BaseStats.SpecialAttack,
				model.SwShPokemonRelation.BaseStats.SpecialDefense,
				model.SwShPokemonRelation.BaseStats.Speed,
				model.SwShPokemonRelation.BaseStats.Average,
				model.SwShPokemonRelation.BaseStats.Total,
			),
			entity.NewIndividualValues(
				model.IndividualValuesId,
				model.Id,
				model.IndividualValues.HitPoints,
				model.IndividualValues.Attack,
				model.IndividualValues.Defense,
				model.IndividualValues.SpecialAttack,
				model.IndividualValues.SpecialDefense,
				model.IndividualValues.Speed,
			),
			entity.NewBasePoints(
				model.BasePointsId,
				model.Id,
				model.BasePoints.HitPoints,
				model.BasePoints.Attack,
				model.BasePoints.Defense,
				model.BasePoints.SpecialAttack,
				model.BasePoints.SpecialDefense,
				model.BasePoints.Speed,
			),
			entity.NewActualValues(
				model.ActualValuesId,
				model.Id,
				model.ActualValues.HitPoints,
				model.ActualValues.Attack,
				model.ActualValues.Defense,
				model.ActualValues.SpecialAttack,
				model.ActualValues.SpecialDefense,
				model.ActualValues.Speed,
			),
			moves,
			model.Note,
		),
		model.IsGigantamax,
	)
}

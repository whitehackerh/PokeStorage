package service

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVBredPokemonService interface {
		MakeEntityFromApiSchema(api_schema.SVBredPokemon, string) entity.SVBredPokemon
		MakeEntityFromModel(model.SVBredPokemonRelation) entity.SVBredPokemon
	}
	SVBredPokemonService struct {
		bredPokemonService IBredPokemonService
	}
)

func NewSVBredPokemonService(bredPokemonService IBredPokemonService) ISVBredPokemonService {
	return &SVBredPokemonService{
		bredPokemonService: bredPokemonService,
	}
}

func (s *SVBredPokemonService) MakeEntityFromApiSchema(schema api_schema.SVBredPokemon, userId string) entity.SVBredPokemon {
	return entity.NewSVBredPokemon(
		s.bredPokemonService.MakeEntityFromApiSchema(schema.BredPokemon, userId),
		entity.NewTeraType(
			schema.TeraType.Id,
			schema.TeraType.Name,
		),
	)
}

func (s *SVBredPokemonService) MakeEntityFromModel(model model.SVBredPokemonRelation) entity.SVBredPokemon {
	var types []entity.Type
	types = append(types, entity.NewType(
		model.SVPokemonRelation.Type1.Id,
		model.SVPokemonRelation.Type1.Name,
	))
	if model.SVPokemonRelation.Type2 != nil {
		types = append(types, entity.NewType(
			model.SVPokemonRelation.Type2.Id,
			model.SVPokemonRelation.Type2.Name,
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

	return entity.NewSVBredPokemon(
		entity.NewBredPokemon(
			model.Id,
			model.UserId,
			model.PokemonId,
			model.SVPokemonRelation.NationalPokedexNo,
			model.SVPokemonRelation.FormeNo,
			model.SVPokemonRelation.Name,
			model.SVPokemonRelation.FormeName,
			entity.NewGender(model.Gender.Id, model.Gender.Name),
			model.Level,
			types,
			entity.NewAbility(model.Ability.Id, model.Ability.Name),
			entity.NewNature(model.Nature.Id, model.Nature.Name),
			heldItem,
			entity.NewBaseStats(
				model.SVPokemonRelation.BaseStats.Id,
				model.PokemonId,
				model.SVPokemonRelation.BaseStats.HitPoints,
				model.SVPokemonRelation.BaseStats.Attack,
				model.SVPokemonRelation.BaseStats.Defense,
				model.SVPokemonRelation.BaseStats.SpecialAttack,
				model.SVPokemonRelation.BaseStats.SpecialDefense,
				model.SVPokemonRelation.BaseStats.Speed,
				model.SVPokemonRelation.BaseStats.Average,
				model.SVPokemonRelation.BaseStats.Total,
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
		entity.NewTeraType(model.TeraType.Id, model.TeraType.Name),
	)
}

package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISwShMoveService interface {
		Make([]model.SwShMoveRelation) []entity.Move
	}
	SwShMoveService struct{}
)

func NewSwShMoveService() ISwShMoveService {
	return &SwShMoveService{}
}

func (s *SwShMoveService) Make(moves []model.SwShMoveRelation) []entity.Move {
	var result []entity.Move
	for _, move := range moves {
		moveType := entity.NewType(
			move.Type.Id,
			move.Type.Name,
			move.Type.FirstAppearanceTitleId,
		)

		moveCategory := entity.NewMoveCategory(
			move.MoveCategory.Id,
			move.MoveCategory.Name,
		)

		result = append(result, entity.NewMove(
			move.Id,
			move.Name,
			moveType,
			moveCategory,
			move.Power,
			move.Accuracy,
		))
	}
	return result
}

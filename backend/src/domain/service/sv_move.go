package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVMoveService interface {
		Make([]model.SVMoveRelation) []entity.Move
	}
	SVMoveService struct{}
)

func NewSVMoveService() ISVMoveService {
	return &SVMoveService{}
}

func (s *SVMoveService) Make(moves []model.SVMoveRelation) []entity.Move {
	var result []entity.Move
	for _, move := range moves {
		moveType := entity.NewType(
			move.Type.Id,
			move.Type.Name,
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

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetMovesPresenter struct{}

func NewGetMovesPresenter() usecase.GetMovesPresenter {
	return &GetMovesPresenter{}
}

func (p *GetMovesPresenter) Output(moves []entity.Move) usecase.GetMovesOutput {
	output := usecase.GetMovesOutput{
		Moves: []api_schema.Move{},
	}
	for _, move := range moves {
		moveTypeEntity := move.MoveType()
		moveType := api_schema.Type{
			Id:   moveTypeEntity.Id(),
			Name: moveTypeEntity.Name(),
		}

		moveCategoryEntity := move.MoveCategory()
		moveCategory := api_schema.MoveCategory{
			Id:   moveCategoryEntity.Id(),
			Name: moveCategoryEntity.Name(),
		}

		output.Moves = append(output.Moves, api_schema.Move{
			Id:           move.Id(),
			Name:         move.Name(),
			MoveType:     moveType,
			MoveCategory: moveCategory,
			Power:        move.Power(),
			Accuracy:     move.Accuracy(),
		})
	}
	return output
}

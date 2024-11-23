package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetSVMovesPresenter struct{}

func NewGetSVMovesPresenter() usecase.GetSVMovesPresenter {
	return &GetSVMovesPresenter{}
}

func (p *GetSVMovesPresenter) Output(moves []entity.Move) usecase.GetSVMovesOutput {
	output := usecase.GetSVMovesOutput{
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

package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetSVTeamsPresenter struct{}

func NewGetSVTeamsPresenter() usecase.GetSVTeamsPresenter {
	return &GetSVTeamsPresenter{}
}

func (p *GetSVTeamsPresenter) Output(teams []entity.SVTeam) usecase.GetSVTeamsOutput {
	output := usecase.GetSVTeamsOutput{
		Teams: []api_schema.SVTeam{},
	}

	for _, team := range teams {
		output.Teams = append(output.Teams, api_schema.SVTeam{
			Id:           team.Id(),
			Name:         team.Name(),
			BredPokemons: p.OutputBredPokemons(team.BredPokemons()),
			Note:         team.Note(),
		})
	}

	return output
}

func (p *GetSVTeamsPresenter) OutputBredPokemons(bredPokemons []*entity.SVBredPokemon) []*api_schema.SVBredPokemon {
	output := []*api_schema.SVBredPokemon{}
	for _, bredPokemon := range bredPokemons {
		if bredPokemon == nil {
			output = append(output, nil)
			continue
		}

		teraType := bredPokemon.TeraType()
		gender := bredPokemon.Gender()
		var types []api_schema.Type
		for _, t := range bredPokemon.Types() {
			types = append(types, api_schema.Type{
				Id:   t.Id(),
				Name: t.Name(),
			})
		}
		ability := bredPokemon.Ability()
		nature := bredPokemon.Nature()
		heldItem := bredPokemon.HeldItem()
		var heldItemResponse *api_schema.Item
		if heldItem == nil {
			heldItemResponse = nil
		} else {
			heldItemResponse = &api_schema.Item{
				Id:   heldItem.Id(),
				Name: heldItem.Name(),
			}
		}
		baseStats := bredPokemon.BaseStats()
		individualValues := bredPokemon.IndividualValues()
		basePoints := bredPokemon.BasePoints()
		actualValues := bredPokemon.ActualValues()
		var moves []api_schema.Move
		for _, m := range bredPokemon.Moves() {
			moveType := m.MoveType()
			moveCategory := m.MoveCategory()
			moves = append(moves, api_schema.Move{
				Id:   m.Id(),
				Name: m.Name(),
				MoveType: api_schema.Type{
					Id:   moveType.Id(),
					Name: moveType.Name(),
				},
				MoveCategory: api_schema.MoveCategory{
					Id:   moveCategory.Id(),
					Name: moveCategory.Name(),
				},
				Power:    m.Power(),
				Accuracy: m.Accuracy(),
			})
		}

		output = append(output, &api_schema.SVBredPokemon{
			BredPokemon: api_schema.BredPokemon{
				Id:                bredPokemon.Id(),
				PokemonId:         bredPokemon.PokemonId(),
				NationalPokedexNo: bredPokemon.NationalPokedexNo(),
				FormeNo:           bredPokemon.FormeNo(),
				Name:              bredPokemon.Name(),
				FormeName:         bredPokemon.FormeName(),
				Gender: api_schema.Gender{
					Id:   gender.Id(),
					Name: gender.Name(),
				},
				Level: bredPokemon.Level(),
				Types: types,
				Ability: api_schema.Ability{
					Id:   ability.Id(),
					Name: ability.Name(),
				},
				Nature: api_schema.Nature{
					Id:   nature.Id(),
					Name: nature.Name(),
				},
				HeldItem: heldItemResponse,
				BaseStats: api_schema.BaseStats{
					Id:             baseStats.Id(),
					HitPoints:      baseStats.HitPoints(),
					Attack:         baseStats.Attack(),
					Defense:        baseStats.Defense(),
					SpecialAttack:  baseStats.SpecialAttack(),
					SpecialDefense: baseStats.SpecialDefense(),
					Speed:          baseStats.Speed(),
					Average:        baseStats.Average(),
					Total:          baseStats.Total(),
				},
				IndividualValues: api_schema.IndividualValues{
					Id:             individualValues.Id(),
					HitPoints:      individualValues.HitPoints(),
					Attack:         individualValues.Attack(),
					Defense:        individualValues.Defense(),
					SpecialAttack:  individualValues.SpecialAttack(),
					SpecialDefense: individualValues.SpecialDefense(),
					Speed:          individualValues.Speed(),
				},
				BasePoints: api_schema.BasePoints{
					Id:             basePoints.Id(),
					HitPoints:      basePoints.HitPoints(),
					Attack:         basePoints.Attack(),
					Defense:        basePoints.Defense(),
					SpecialAttack:  basePoints.SpecialAttack(),
					SpecialDefense: basePoints.SpecialDefense(),
					Speed:          basePoints.Speed(),
				},
				ActualValues: api_schema.ActualValues{
					Id:             actualValues.Id(),
					HitPoints:      actualValues.HitPoints(),
					Attack:         actualValues.Attack(),
					Defense:        actualValues.Defense(),
					SpecialAttack:  actualValues.SpecialAttack(),
					SpecialDefense: actualValues.SpecialDefense(),
					Speed:          actualValues.Speed(),
				},
				Moves: moves,
				Note:  bredPokemon.Note(),
			},
			TeraType: api_schema.TeraType{
				Id:   teraType.Id(),
				Name: teraType.Name(),
			},
		})
	}

	return output
}

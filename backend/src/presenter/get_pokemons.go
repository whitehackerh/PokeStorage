package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/response_component"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetPokemonsPresenter struct{}

func NewGetPokemonsPresenter() usecase.GetPokemonsPresenter {
	return &GetPokemonsPresenter{}
}

func (p *GetPokemonsPresenter) Output(pokemons []entity.Pokemon) usecase.GetPokemonsOutput {
	output := usecase.GetPokemonsOutput{
		Pokemons: make([]response_component.Pokemon, len(pokemons)),
	}
	for i, pokemon := range pokemons {
		types := make([]response_component.Type, len(pokemon.Types()))
		for j, typeEntity := range pokemon.Types() {
			types[j] = response_component.Type{
				Id:   typeEntity.Id(),
				Name: typeEntity.Name(),
			}
		}

		abilities := make([]response_component.Ability, len(pokemon.Abilities()))
		for j, ability := range pokemon.Abilities() {
			abilities[j] = response_component.Ability{
				Id:   ability.Id(),
				Name: ability.Name(),
			}
		}

		baseStats := pokemon.BaseStats()
		baseStatsResponse := response_component.BaseStats{
			Id:             baseStats.Id(),
			HitPoints:      baseStats.HitPoints(),
			Attack:         baseStats.Attack(),
			Defense:        baseStats.Defense(),
			SpecialAttack:  baseStats.SpecialAttack(),
			SpecialDefense: baseStats.SpecialDefense(),
			Speed:          baseStats.Speed(),
			Average:        baseStats.Average(),
			Total:          baseStats.Total(),
		}

		presetHeldItem := pokemon.PresetHeldItem()
		var presetHeldItemResponse *response_component.Item
		if presetHeldItem == nil {
			presetHeldItemResponse = nil
		} else {
			presetHeldItemResponse = &response_component.Item{
				Id:   presetHeldItem.Id(),
				Name: presetHeldItem.Name(),
			}
		}

		output.Pokemons[i] = response_component.Pokemon{
			Id:                pokemon.Id(),
			NationalPokedexNo: pokemon.NationalPokedexNo(),
			FormeNo:           pokemon.FormeNo(),
			Name:              pokemon.Name(),
			FormeName:         pokemon.FormeName(),
			Types:             types,
			Abilities:         abilities,
			BaseStats:         baseStatsResponse,
			PresetHeldItem:    presetHeldItemResponse,
		}
	}
	return output
}

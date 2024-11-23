package presenter

import (
	"github.com/whitehackerh/PokeStorage/src/api_schema"
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

type GetSwShPokemonsPresenter struct{}

func NewGetSwShPokemonsPresenter() usecase.GetSwShPokemonsPresenter {
	return &GetSwShPokemonsPresenter{}
}

func (p *GetSwShPokemonsPresenter) Output(pokemons []entity.Pokemon) usecase.GetSwShPokemonsOutput {
	output := usecase.GetSwShPokemonsOutput{
		Pokemons: make([]api_schema.Pokemon, len(pokemons)),
	}
	for i, pokemon := range pokemons {
		types := make([]api_schema.Type, len(pokemon.Types()))
		for j, typeEntity := range pokemon.Types() {
			types[j] = api_schema.Type{
				Id:   typeEntity.Id(),
				Name: typeEntity.Name(),
			}
		}

		abilities := make([]api_schema.Ability, len(pokemon.Abilities()))
		for j, ability := range pokemon.Abilities() {
			abilities[j] = api_schema.Ability{
				Id:   ability.Id(),
				Name: ability.Name(),
			}
		}

		baseStats := pokemon.BaseStats()
		baseStatsResponse := api_schema.BaseStats{
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
		var presetHeldItemResponse *api_schema.Item
		if presetHeldItem == nil {
			presetHeldItemResponse = nil
		} else {
			presetHeldItemResponse = &api_schema.Item{
				Id:   presetHeldItem.Id(),
				Name: presetHeldItem.Name(),
			}
		}

		output.Pokemons[i] = api_schema.Pokemon{
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

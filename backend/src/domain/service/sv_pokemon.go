package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	ISVPokemonService interface {
		Make([]model.SVPokemonRelation) []entity.Pokemon
	}
	SVPokemonService struct{}
)

func NewSVPokemonService() ISVPokemonService {
	return &SVPokemonService{}
}

func (s *SVPokemonService) Make(pokemons []model.SVPokemonRelation) []entity.Pokemon {
	var result []entity.Pokemon
	for _, pokemon := range pokemons {
		var types []entity.Type
		types = append(types, entity.NewType(
			pokemon.Type1.Id,
			pokemon.Type1.Name,
		))
		if pokemon.Type2 != nil {
			types = append(types, entity.NewType(
				pokemon.Type2.Id,
				pokemon.Type2.Name,
			))
		}

		var abilities []entity.Ability
		abilities = append(abilities, entity.NewAbility(
			pokemon.Ability1.Id,
			pokemon.Ability1.Name,
		))
		if pokemon.Ability2 != nil {
			abilities = append(abilities, entity.NewAbility(
				pokemon.Ability2.Id,
				pokemon.Ability2.Name,
			))
		}
		if pokemon.HiddenAbility != nil {
			abilities = append(abilities, entity.NewAbility(
				pokemon.HiddenAbility.Id,
				pokemon.HiddenAbility.Name,
			))
		}

		baseStats := entity.NewBaseStats(
			pokemon.BaseStats.Id,
			pokemon.BaseStats.PokemonId,
			pokemon.BaseStats.HitPoints,
			pokemon.BaseStats.Attack,
			pokemon.BaseStats.Defense,
			pokemon.BaseStats.SpecialAttack,
			pokemon.BaseStats.SpecialDefense,
			pokemon.BaseStats.Speed,
			pokemon.BaseStats.Average,
			pokemon.BaseStats.Total,
		)

		var presetHeldItem *entity.Item
		if pokemon.PresetHeldItem != nil {
			item := entity.NewItem(
				pokemon.PresetHeldItem.Id,
				pokemon.PresetHeldItem.Name,
			)
			presetHeldItem = &item
		}

		result = append(result, entity.NewPokemon(
			pokemon.Id,
			pokemon.NationalPokedexNo,
			pokemon.FormeNo,
			pokemon.Name,
			pokemon.FormeName,
			types,
			abilities,
			baseStats,
			presetHeldItem,
		))
	}
	return result
}

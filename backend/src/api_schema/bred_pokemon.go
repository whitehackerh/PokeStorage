package api_schema

type BredPokemon struct {
	Id                string           `json:"id"`
	PokemonId         int              `json:"pokemon_id"`
	NationalPokedexNo int              `json:"national_pokedex_no"`
	FormeNo           int              `json:"forme_no"`
	Name              string           `json:"name"`
	FormeName         *string          `json:"forme_name"`
	Gender            Gender           `json:"gender"`
	Level             int              `json:"level"`
	Types             []Type           `json:"types"`
	Ability           Ability          `json:"ability"`
	Nature            Nature           `json:"nature"`
	HeldItem          *Item            `json:"held_item"`
	BaseStats         BaseStats        `json:"base_stats"`
	IndividualValues  IndividualValues `json:"individual_values"`
	BasePoints        BasePoints       `json:"base_points"`
	ActualValues      ActualValues     `json:"actual_values"`
	Moves             []Move           `json:"moves"`
	Note              *string          `json:"note"`
}

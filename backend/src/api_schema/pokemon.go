package api_schema

type Pokemon struct {
	Id                int       `json:"id"`
	NationalPokedexNo int       `json:"national_pokedex_no"`
	FormeNo           int       `json:"forme_no"`
	Name              string    `json:"name"`
	FormeName         *string   `json:"forme_name"`
	Types             []Type    `json:"types"`
	Abilities         []Ability `json:"abilities"`
	BaseStats         BaseStats `json:"base_stats"`
	PresetHeldItem    *Item     `json:"preset_held_item"`
}

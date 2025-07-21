package api_schema

type SVTeam struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	BredPokemons []*SVBredPokemon `json:"bred_pokemons"`
	Note         *string          `json:"note"`
}

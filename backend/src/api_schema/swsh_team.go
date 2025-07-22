package api_schema

type SwShTeam struct {
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	BredPokemons []*SwShBredPokemon `json:"bred_pokemons"`
	Note         *string            `json:"note"`
}

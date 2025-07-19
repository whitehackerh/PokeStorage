package api_schema

type PostPutTeam struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	BredPokemonIds []*string `json:"bred_pokemon_ids"`
	Note           *string   `json:"note"`
}

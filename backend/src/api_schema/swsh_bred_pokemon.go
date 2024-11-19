package api_schema

type SwShBredPokemon struct {
	BredPokemon
	IsGigantamax bool `json:"is_gigantamax"`
}

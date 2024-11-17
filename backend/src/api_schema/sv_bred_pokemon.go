package api_schema

type SVBredPokemon struct {
	BredPokemon
	TeraType TeraType `json:"tera_type"`
}

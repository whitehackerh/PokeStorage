package entity

type SwShTeam struct {
	Team
	bredPokemons []*SwShBredPokemon
}

func NewSwShTeam(
	team Team,
	bredPokemons []*SwShBredPokemon,
) SwShTeam {
	return SwShTeam{
		team,
		bredPokemons,
	}
}

func (st *SwShTeam) BredPokemons() []*SwShBredPokemon {
	return st.bredPokemons
}

package entity

type SVTeam struct {
	Team
	bredPokemons []SVBredPokemon
}

func NewSVTeam(
	team Team,
	bredPokemons []SVBredPokemon,
) SVTeam {
	return SVTeam{
		team,
		bredPokemons,
	}
}

func (st *SVTeam) BredPokemons() []SVBredPokemon {
	return st.bredPokemons
}

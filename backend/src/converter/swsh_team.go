package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
	"github.com/whitehackerh/PokeStorage/src/util"
)

func SwShTeamEntityToModel(team entity.SwShTeam) model.SwShTeam {
	bredPokemonIds := []*string{nil, nil, nil, nil, nil, nil}
	for i, bredPokemon := range team.BredPokemons() {
		bredPokemonIds[i] = util.StringToPointer(bredPokemon.Id())
	}

	return model.SwShTeam{
		Id:             team.Id(),
		UserId:         team.UserId(),
		Name:           team.Name(),
		BredPokemon1Id: bredPokemonIds[0],
		BredPokemon2Id: bredPokemonIds[1],
		BredPokemon3Id: bredPokemonIds[2],
		BredPokemon4Id: bredPokemonIds[3],
		BredPokemon5Id: bredPokemonIds[4],
		BredPokemon6Id: bredPokemonIds[5],
		Note:           team.Note(),
	}
}

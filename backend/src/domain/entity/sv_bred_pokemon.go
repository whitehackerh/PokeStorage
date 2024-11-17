package entity

type SVBredPokemon struct {
	BredPokemon
	teraType TeraType
}

func NewSVBredPokemon(
	bredPokemon BredPokemon,
	teraType TeraType,
) SVBredPokemon {
	return SVBredPokemon{
		bredPokemon,
		teraType,
	}
}

func (svb *SVBredPokemon) TeraType() TeraType {
	return svb.teraType
}

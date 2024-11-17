package entity

type SwShBredPokemon struct {
	BredPokemon
	isGigantamax bool
}

func NewSwShBredPokemon(
	bredPokemon BredPokemon,
	isGigantamax bool,
) SwShBredPokemon {
	return SwShBredPokemon{
		bredPokemon,
		isGigantamax,
	}
}

func (ssb *SwShBredPokemon) IsGigantamax() bool {
	return ssb.isGigantamax
}

package entity

type IndividualValues struct {
	id             string
	bredPokemonId  string
	hitPoints      int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

func NewIndividualValues(
	id string,
	bredPokemonId string,
	hitPoints int,
	attack int,
	defense int,
	specialAttack int,
	specialDefense int,
	speed int,
) IndividualValues {
	return IndividualValues{
		id:             id,
		bredPokemonId:  bredPokemonId,
		hitPoints:      hitPoints,
		attack:         attack,
		defense:        defense,
		specialAttack:  specialAttack,
		specialDefense: specialDefense,
		speed:          speed,
	}
}

func (i *IndividualValues) Id() string {
	return i.id
}

func (i *IndividualValues) BredPokemonId() string {
	return i.bredPokemonId
}

func (i *IndividualValues) HitPoints() int {
	return i.hitPoints
}

func (i *IndividualValues) Attack() int {
	return i.attack
}

func (i *IndividualValues) Defense() int {
	return i.defense
}

func (i *IndividualValues) SpecialAttack() int {
	return i.specialAttack
}

func (i *IndividualValues) SpecialDefense() int {
	return i.specialDefense
}

func (i *IndividualValues) Speed() int {
	return i.speed
}

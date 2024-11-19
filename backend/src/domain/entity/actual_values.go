package entity

type ActualValues struct {
	id             string
	bredPokemonId  string
	hitPoints      int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

func NewActualValues(
	id string,
	bredPokemonId string,
	hitPoints int,
	attack int,
	defense int,
	specialAttack int,
	specialDefense int,
	speed int,
) ActualValues {
	return ActualValues{
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

func (a *ActualValues) Id() string {
	return a.id
}

func (a *ActualValues) BredPokemonId() string {
	return a.bredPokemonId
}

func (a *ActualValues) HitPoints() int {
	return a.hitPoints
}

func (a *ActualValues) Attack() int {
	return a.attack
}

func (a *ActualValues) Defense() int {
	return a.defense
}

func (a *ActualValues) SpecialAttack() int {
	return a.specialAttack
}

func (a *ActualValues) SpecialDefense() int {
	return a.specialDefense
}

func (a *ActualValues) Speed() int {
	return a.speed
}

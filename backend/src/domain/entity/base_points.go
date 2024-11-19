package entity

type BasePoints struct {
	id             string
	bredPokemonId  string
	hitPoints      int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

func NewBasePoints(
	id string,
	bredPokemonId string,
	hitPoints int,
	attack int,
	defense int,
	specialAttack int,
	specialDefense int,
	speed int,
) BasePoints {
	return BasePoints{
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

func (e *BasePoints) Id() string {
	return e.id
}

func (e *BasePoints) BredPokemonId() string {
	return e.bredPokemonId
}

func (e *BasePoints) HitPoints() int {
	return e.hitPoints
}

func (e *BasePoints) Attack() int {
	return e.attack
}

func (e *BasePoints) Defense() int {
	return e.defense
}

func (e *BasePoints) SpecialAttack() int {
	return e.specialAttack
}

func (e *BasePoints) SpecialDefense() int {
	return e.specialDefense
}

func (e *BasePoints) Speed() int {
	return e.speed
}

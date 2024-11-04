package entity

type BaseStats struct {
	id             int
	pokemonId      int
	hitPoints      int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
	average        float64
	total          int
}

func NewBaseStats(
	id int,
	pokemonId int,
	hitPoints int,
	attack int,
	defense int,
	specialAttack int,
	specialDefense int,
	speed int,
	average float64,
	total int,
) BaseStats {
	return BaseStats{
		id:             id,
		pokemonId:      pokemonId,
		hitPoints:      hitPoints,
		attack:         attack,
		defense:        defense,
		specialAttack:  specialAttack,
		specialDefense: specialDefense,
		speed:          speed,
		average:        average,
		total:          total,
	}
}

func (b *BaseStats) Id() int {
	return b.id
}

func (b *BaseStats) PokemonId() int {
	return b.pokemonId
}

func (b *BaseStats) HitPoints() int {
	return b.hitPoints
}

func (b *BaseStats) Attack() int {
	return b.attack
}

func (b *BaseStats) Defense() int {
	return b.defense
}

func (b *BaseStats) SpecialAttack() int {
	return b.specialAttack
}

func (b *BaseStats) SpecialDefense() int {
	return b.specialDefense
}

func (b *BaseStats) Speed() int {
	return b.speed
}

func (b *BaseStats) Average() float64 {
	return b.average
}

func (b *BaseStats) Total() int {
	return b.total
}

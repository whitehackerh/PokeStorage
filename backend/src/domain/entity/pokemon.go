package entity

type Pokemon struct {
	id                int
	nationalPokedexNo int
	formeNo           int
	name              string
	formeName         string
	types             []Type
	abilities         []Ability
	baseStats         BaseStats
	presetHeldItem    *Item
}

func NewPokemon(
	id int,
	nationalPokedexNo int,
	formeNo int,
	name string,
	formeName string,
	types []Type,
	abilities []Ability,
	baseStats BaseStats,
	presetHeldItem *Item,
) Pokemon {
	return Pokemon{
		id:                id,
		nationalPokedexNo: nationalPokedexNo,
		formeNo:           formeNo,
		name:              name,
		formeName:         formeName,
		types:             types,
		abilities:         abilities,
		baseStats:         baseStats,
		presetHeldItem:    presetHeldItem,
	}
}

func (p *Pokemon) Id() int {
	return p.id
}

func (p *Pokemon) NationalPokedexNo() int {
	return p.nationalPokedexNo
}

func (p *Pokemon) FormeNo() int {
	return p.formeNo
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) FormeName() string {
	return p.formeName
}

func (p *Pokemon) Types() []Type {
	return p.types
}

func (p *Pokemon) Abilities() []Ability {
	return p.abilities
}

func (p *Pokemon) BaseStats() BaseStats {
	return p.baseStats
}

func (p *Pokemon) PresetHeldItem() *Item {
	return p.presetHeldItem
}

package entity

type BredPokemon struct {
	id                string
	userId            string
	pokemonId         int
	nationalPokedexNo int
	formeNo           int
	name              string
	formeName         *string
	gender            Gender
	level             int
	types             []Type
	ability           Ability
	nature            Nature
	heldItem          *Item
	baseStats         BaseStats
	individualValues  IndividualValues
	basePoints        BasePoints
	actualValues      ActualValues
	moves             []Move
	note              *string
}

func NewBredPokemon(
	id string,
	userId string,
	pokemonId int,
	nationalPokedexNo int,
	formeNo int,
	name string,
	formeName *string,
	gender Gender,
	level int,
	types []Type,
	ability Ability,
	nature Nature,
	heldItem *Item,
	baseStats BaseStats,
	individualValues IndividualValues,
	basePoints BasePoints,
	actualValues ActualValues,
	moves []Move,
	note *string,
) BredPokemon {
	return BredPokemon{
		id:                id,
		userId:            userId,
		pokemonId:         pokemonId,
		nationalPokedexNo: nationalPokedexNo,
		formeNo:           formeNo,
		name:              name,
		formeName:         formeName,
		gender:            gender,
		level:             level,
		types:             types,
		ability:           ability,
		nature:            nature,
		heldItem:          heldItem,
		baseStats:         baseStats,
		individualValues:  individualValues,
		basePoints:        basePoints,
		actualValues:      actualValues,
		moves:             moves,
		note:              note,
	}
}

func (bp *BredPokemon) Id() string {
	return bp.id
}

func (bp *BredPokemon) UserId() string {
	return bp.userId
}

func (bp *BredPokemon) PokemonId() int {
	return bp.pokemonId
}

func (bp *BredPokemon) NationalPokedexNo() int {
	return bp.nationalPokedexNo
}

func (bp *BredPokemon) FormeNo() int {
	return bp.formeNo
}

func (bp *BredPokemon) Name() string {
	return bp.name
}

func (bp *BredPokemon) FormeName() *string {
	return bp.formeName
}

func (bp *BredPokemon) Gender() Gender {
	return bp.gender
}

func (bp *BredPokemon) Level() int {
	return bp.level
}

func (bp *BredPokemon) Types() []Type {
	return bp.types
}

func (bp *BredPokemon) Ability() Ability {
	return bp.ability
}

func (bp *BredPokemon) Nature() Nature {
	return bp.nature
}

func (bp *BredPokemon) HeldItem() *Item {
	return bp.heldItem
}

func (bp *BredPokemon) BaseStats() BaseStats {
	return bp.baseStats
}

func (bp *BredPokemon) IndividualValues() IndividualValues {
	return bp.individualValues
}

func (bp *BredPokemon) BasePoints() BasePoints {
	return bp.basePoints
}

func (bp *BredPokemon) ActualValues() ActualValues {
	return bp.actualValues
}

func (bp *BredPokemon) Moves() []Move {
	return bp.moves
}

func (bp *BredPokemon) Note() *string {
	return bp.note
}

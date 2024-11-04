package model

type SVPokemon struct {
	Id                int    `gorm:"primaryKey"`
	NationalPokedexNo int    `gorm:"column:national_pokedex_no;not null"`
	FormeNo           int    `gorm:"column:forme_no;primaryKey;not null"`
	Name              string `gorm:"column:name;not null"`
	FormeName         string `gorm:"column:forme_name;not null"`
	Type1Id           int    `gorm:"column:type_1_id;not null"`
	Type2Id           *int   `gorm:"column:type_2_id"`
	Ability1Id        int    `gorm:"column:ability_1_id;not null"`
	Ability2Id        *int   `gorm:"column:ability_2_id"`
	HiddenAbilityId   *int   `gorm:"column:hidden_ability_id"`
	BaseStatsId       int    `gorm:"column:base_stats_id;not null"`
	PresetHeldItemId  *int   `gorm:"column:preset_held_item_id"`
}

type SVPokemonRelation struct {
	SVPokemon
	Type1          Type               `gorm:"foreignKey:Type1Id;references:Id"`
	Type2          *Type              `gorm:"foreignKey:Type2Id;references:Id"`
	Ability1       Ability            `gorm:"foreignKey:Ability1Id;references:Id"`
	Ability2       *Ability           `gorm:"foreignKey:Ability2Id;references:Id"`
	HiddenAbility  *Ability           `gorm:"foreignKey:HiddenAbilityId;references:Id"`
	BaseStats      SVPokemonBaseStats `gorm:"foreignKey:Id;references:PokemonId"`
	PresetHeldItem *SVItem            `gorm:"foreignKey:PresetHeldItemId;references:Id"`
}

func (m SVPokemonRelation) TableName() string {
	return "sv_pokemons"
}

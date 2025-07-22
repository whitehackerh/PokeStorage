package model

import "time"

type SwShTeam struct {
	Id             string     `gorm:"primaryKey"`
	UserId         string     `gorm:"column:user_id;not null"`
	Name           string     `gorm:"column:name;not null"`
	BredPokemon1Id *string    `gorm:"column:bred_pokemon_1_id"`
	BredPokemon2Id *string    `gorm:"column:bred_pokemon_2_id"`
	BredPokemon3Id *string    `gorm:"column:bred_pokemon_3_id"`
	BredPokemon4Id *string    `gorm:"column:bred_pokemon_4_id"`
	BredPokemon5Id *string    `gorm:"column:bred_pokemon_5_id"`
	BredPokemon6Id *string    `gorm:"column:bred_pokemon_6_id"`
	Note           *string    `gorm:"column:note"`
	CreatedAt      *time.Time `gorm:"autoCreateTime"`
	UpdatedAt      *time.Time `gorm:"autoUpdateTime"`
	DeletedAt      *time.Time `gorm:"index"`
}

type SwShTeamRelation struct {
	SwShTeam
	BredPokemon1 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon1Id;references:Id"`
	BredPokemon2 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon2Id;references:Id"`
	BredPokemon3 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon3Id;references:Id"`
	BredPokemon4 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon4Id;references:Id"`
	BredPokemon5 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon5Id;references:Id"`
	BredPokemon6 *SwShBredPokemonRelation `gorm:"foreignKey:BredPokemon6Id;references:Id"`
}

func (s *SwShTeamRelation) TableName() string {
	return "swsh_teams"
}

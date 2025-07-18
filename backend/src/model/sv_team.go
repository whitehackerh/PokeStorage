package model

import "time"

type SVTeam struct {
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

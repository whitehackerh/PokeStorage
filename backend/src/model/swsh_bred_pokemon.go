package model

import "time"

type SwShBredPokemon struct {
	Id                 string     `gorm:"primaryKey"`
	UserId             string     `gorm:"column:user_id;not null"`
	PokemonId          int        `gorm:"column:pokemon_id;not null"`
	GenderId           int        `gorm:"column:gender_id;not null"`
	Level              int        `gorm:"column:level;not null"`
	IsGigantamax       bool       `gorm:"column:is_gigantamax;not null"`
	AbilityId          int        `gorm:"column:ability_id;not null"`
	NatureId           int        `gorm:"column:nature_id;not null"`
	ItemId             *int       `gorm:"column:item_id"`
	IndividualValuesId string     `gorm:"column:individual_values_id;not null"`
	BasePointsId       string     `gorm:"column:base_points_id;not null"`
	ActualValuesId     string     `gorm:"column:actual_values_id;not null"`
	Move1Id            int        `gorm:"column:move_1_id;not null"`
	Move2Id            *int       `gorm:"column:move_2_id"`
	Move3Id            *int       `gorm:"column:move_3_id"`
	Move4Id            *int       `gorm:"column:move_4_id"`
	Note               *string    `gorm:"column:note"`
	CreatedAt          *time.Time `gorm:"autoCreateTime"`
	UpdatedAt          *time.Time `gorm:"autoUpdateTime"`
	DeletedAt          *time.Time `gorm:"index"`
}

type SwShBredPokemonRelation struct {
	SwShBredPokemon
	SwShPokemonRelation SwShPokemonRelation  `gorm:"foreignKey:PokemonId;references:Id"`
	Gender              Gender               `gorm:"foreignKey:GenderId;references:Id"`
	Ability             Ability              `gorm:"foreignKey:AbilityId;references:Id"`
	Nature              Nature               `gorm:"foreignKey:NatureId;references:Id"`
	HeldItem            *SwShItem            `gorm:"foreignKey:ItemId;references:Id"`
	IndividualValues    SwShIndividualValues `gorm:"foreignKey:Id;references:BredPokemonId"`
	BasePoints          SwShBasePoints       `gorm:"foreignKey:Id;references:BredPokemonId"`
	ActualValues        SwShActualValues     `gorm:"foreignKey:Id;references:BredPokemonId"`
	Move1Relation       SwShMoveRelation     `gorm:"foreignKey:Move1Id;references:Id"`
	Move2Relation       *SwShMoveRelation    `gorm:"foreignKey:Move2Id;references:Id"`
	Move3Relation       *SwShMoveRelation    `gorm:"foreignKey:Move3Id;references:Id"`
	Move4Relation       *SwShMoveRelation    `gorm:"foreignKey:Move4Id;references:Id"`
}

func (s *SwShBredPokemonRelation) TableName() string {
	return "swsh_bred_pokemons"
}

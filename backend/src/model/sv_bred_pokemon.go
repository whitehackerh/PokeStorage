package model

type SVBredPokemon struct {
	Id                 string  `gorm:"primaryKey"`
	UserId             string  `gorm:"column:user_id;not null"`
	PokemonId          int     `gorm:"column:pokemon_id;not null"`
	GenderId           int     `gorm:"column:gender_id;not null"`
	Level              int     `gorm:"column:level;not null"`
	TeraTypeId         int     `gorm:"column:tera_type_id;not null"`
	AbilityId          int     `gorm:"column:ability_id;not null"`
	NatureId           int     `gorm:"column:nature_id;not null"`
	ItemId             *int    `gorm:"column:item_id"`
	IndividualValuesId string  `gorm:"column:individual_values_id;not null"`
	BasePointsId       string  `gorm:"column:base_points_id;not null"`
	ActualValuesId     string  `gorm:"column:actual_values_id;not null"`
	Move1Id            int     `gorm:"column:move_1_id;not null"`
	Move2Id            *int    `gorm:"column:move_2_id"`
	Move3Id            *int    `gorm:"column:move_3_id"`
	Move4Id            *int    `gorm:"column:move_4_id"`
	Note               *string `gorm:"column:note"`
}

package model

type SVBasePoints struct {
	Id             string `gorm:"primaryKey"`
	BredPokemonId  string `gorm:"column:bred_pokemon_id;not null"`
	HitPoints      int    `gorm:"column:hit_points;not null"`
	Attack         int    `gorm:"column:attack;not null"`
	Defense        int    `gorm:"column:defense;not null"`
	SpecialAttack  int    `gorm:"column:special_attack;not null"`
	SpecialDefense int    `gorm:"column:special_defense;not null"`
	Speed          int    `gorm:"column:speed;not null"`
}

func (s *SVBasePoints) TableName() string {
	return "sv_base_points"
}

package model

type SVPokemonBaseStats struct {
	Id             int     `gorm:"primaryKey"`
	PokemonId      int     `gorm:"column:pokemon_id;not null"`
	HitPoints      int     `gorm:"column:hit_points;not null"`
	Attack         int     `gorm:"column:attack;not null"`
	Defense        int     `gorm:"column:defense;not null"`
	SpecialAttack  int     `gorm:"column:special_attack;not null"`
	SpecialDefense int     `gorm:"column:special_defense;not null"`
	Speed          int     `gorm:"column:speed;not null"`
	Average        float64 `gorm:"column:average;not null"`
	Total          int     `gorm:"column:total;not null"`
}

func (m SVPokemonBaseStats) TableName() string {
	return "sv_pokemon_base_stats"
}

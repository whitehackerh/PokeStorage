package api_schema

type BaseStats struct {
	Id             int     `json:"id"`
	HitPoints      int     `json:"hit_points"`
	Attack         int     `json:"attack"`
	Defense        int     `json:"defense"`
	SpecialAttack  int     `json:"special_attack"`
	SpecialDefense int     `json:"special_defense"`
	Speed          int     `json:"speed"`
	Average        float64 `json:"average"`
	Total          int     `json:"total"`
}

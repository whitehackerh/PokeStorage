package api_schema

type IndividualValues struct {
	Id             string `json:"id"`
	HitPoints      int    `json:"hit_points"`
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	SpecialAttack  int    `json:"special_attack"`
	SpecialDefense int    `json:"special_defense"`
	Speed          int    `json:"speed"`
}

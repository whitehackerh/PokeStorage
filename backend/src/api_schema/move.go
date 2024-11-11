package api_schema

type Move struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	MoveType     Type         `json:"type"`
	MoveCategory MoveCategory `json:"move_category"`
	Power        *int         `json:"power"`
	Accuracy     *int         `json:"accuracy"`
}

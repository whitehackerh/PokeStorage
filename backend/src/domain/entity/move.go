package entity

type Move struct {
	id           int
	name         string
	moveType     Type
	moveCategory MoveCategory
	power        *int
	accuracy     *int
}

func NewMove(
	id int,
	name string,
	moveType Type,
	moveCategory MoveCategory,
	power *int,
	accuracy *int,
) Move {
	return Move{
		id:           id,
		name:         name,
		moveType:     moveType,
		moveCategory: moveCategory,
		power:        power,
		accuracy:     accuracy,
	}
}

func (m *Move) Id() int {
	return m.id
}

func (m *Move) Name() string {
	return m.name
}

func (m *Move) MoveType() Type {
	return m.moveType
}

func (m *Move) MoveCategory() MoveCategory {
	return m.moveCategory
}

func (m *Move) Power() *int {
	return m.power
}

func (m *Move) Accuracy() *int {
	return m.accuracy
}

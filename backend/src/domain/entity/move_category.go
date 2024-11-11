package entity

type MoveCategory struct {
	id   int
	name string
}

func NewMoveCategory(
	id int,
	name string,
) MoveCategory {
	return MoveCategory{
		id:   id,
		name: name,
	}
}

func (m *MoveCategory) Id() int {
	return m.id
}

func (m *MoveCategory) Name() string {
	return m.name
}

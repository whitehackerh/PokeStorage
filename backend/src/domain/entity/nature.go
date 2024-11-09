package entity

type Nature struct {
	id   int
	name string
}

func NewNature(
	id int,
	name string,
) Nature {
	return Nature{
		id:   id,
		name: name,
	}
}

func (t *Nature) Id() int {
	return t.id
}

func (t *Nature) Name() string {
	return t.name
}

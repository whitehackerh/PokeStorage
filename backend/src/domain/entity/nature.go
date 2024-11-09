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

func (n *Nature) Id() int {
	return n.id
}

func (n *Nature) Name() string {
	return n.name
}

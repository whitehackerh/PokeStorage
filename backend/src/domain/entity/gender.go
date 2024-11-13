package entity

type Gender struct {
	id   int
	name string
}

func NewGender(
	id int,
	name string,
) Gender {
	return Gender{
		id:   id,
		name: name,
	}
}

func (g *Gender) Id() int {
	return g.id
}

func (g *Gender) Name() string {
	return g.name
}

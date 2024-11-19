package entity

type Type struct {
	id   int
	name string
}

func NewType(
	id int,
	name string,
) Type {
	return Type{
		id:   id,
		name: name,
	}
}

func (t *Type) Id() int {
	return t.id
}

func (t *Type) Name() string {
	return t.name
}

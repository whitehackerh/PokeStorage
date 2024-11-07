package entity

type TeraType struct {
	id   int
	name string
}

func NewTeraType(
	id int,
	name string,
) TeraType {
	return TeraType{
		id:   id,
		name: name,
	}
}

func (t *TeraType) Id() int {
	return t.id
}

func (t *TeraType) Name() string {
	return t.name
}

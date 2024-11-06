package entity

type Type struct {
	id                     int
	name                   string
	firstAppearanceTitleId int
}

func NewType(
	id int,
	name string,
	firstAppearanceTitleId int,
) Type {
	return Type{
		id:                     id,
		name:                   name,
		firstAppearanceTitleId: firstAppearanceTitleId,
	}
}

func (t *Type) Id() int {
	return t.id
}

func (t *Type) Name() string {
	return t.name
}

func (t *Type) FirstAppearanceTitleId() int {
	return t.firstAppearanceTitleId
}

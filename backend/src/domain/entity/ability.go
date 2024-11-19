package entity

type Ability struct {
	id   int
	name string
}

func NewAbility(
	id int,
	name string,
) Ability {
	return Ability{
		id:   id,
		name: name,
	}
}

func (a *Ability) Id() int {
	return a.id
}

func (a *Ability) Name() string {
	return a.name
}

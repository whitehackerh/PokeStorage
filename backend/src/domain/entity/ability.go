package entity

type Ability struct {
	id                     int
	name                   string
	firstAppearanceTitleId int
}

func NewAbility(
	id int,
	name string,
	firstAppearanceTitleId int,
) Ability {
	return Ability{
		id:                     id,
		name:                   name,
		firstAppearanceTitleId: firstAppearanceTitleId,
	}
}

func (a *Ability) Id() int {
	return a.id
}

func (a *Ability) Name() string {
	return a.name
}

func (a *Ability) FirstAppearanceTitleId() int {
	return a.firstAppearanceTitleId
}

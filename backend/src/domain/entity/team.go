package entity

type Team struct {
	id     string
	userId string
	name   string
	note   *string
}

func NewTeam(
	id string,
	userId string,
	name string,
	note *string,
) Team {
	return Team{
		id:     id,
		userId: userId,
		name:   name,
		note:   note,
	}
}

func (t *Team) Id() string {
	return t.id
}

func (t *Team) UserId() string {
	return t.userId
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) Note() *string {
	return t.note
}

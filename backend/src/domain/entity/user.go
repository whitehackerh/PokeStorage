package entity

type User struct {
	id           string
	username     string
	password     string
	emailAddress string
	name         string
}

func NewUser(
	id string,
	username string,
	password string,
	emailAddress string,
	name string,
) User {
	return User{
		id:           id,
		username:     username,
		password:     password,
		emailAddress: emailAddress,
		name:         name,
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Password() string {
	return u.password
}

func (u *User) EmailAddress() string {
	return u.emailAddress
}

func (u *User) Name() string {
	return u.name
}

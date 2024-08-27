package entity

import (
	"time"
)

type (
	Title struct {
		id          int
		title       string
		releaseDate time.Time
	}
	ITitleService interface {
		Get() ([]Title, error)
	}
)

func NewTitle(id int, title string, releaseDate time.Time) Title {
	return Title{
		id:          id,
		title:       title,
		releaseDate: releaseDate,
	}
}

func (t *Title) Id() int {
	return t.id
}

func (t *Title) Title() string {
	return t.title
}

func (t *Title) ReleaseDate() time.Time {
	return t.releaseDate
}

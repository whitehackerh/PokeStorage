package model

import "time"

type Title struct {
	Id          int       `gorm:"primaryKey"`
	Title       string    `gorm:"column:title;unique;not null"`
	ReleaseDate time.Time `gorm:"column:release_date;not null"`
}

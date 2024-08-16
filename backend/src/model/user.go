package model

import "time"

type User struct {
	Id           string    `gorm:"primaryKey"`
	Username     string    `gorm:"column:username;unique;not null"`
	Password     string    `gorm:"column:password;not null"`
	EmailAddress string    `gorm:"column:email_address;not null"`
	Name         string    `gorm:"column:name;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	DeletedAt    time.Time `gorm:"index"`
}

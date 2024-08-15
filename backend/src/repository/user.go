package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) service.IUserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) Create(user model.User) error {
	result := ur.Db.Create(&user)
	return result.Error
}

package repository

import (
	"gorm.io/gorm"

	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	IUserRepository interface {
		Create(model.User) error
		FindByUsernameAndPassword(string, string) (model.User, error)
	}
	UserRepository struct {
		Db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) Create(user model.User) error {
	result := ur.Db.Create(&user)
	return result.Error
}

func (ur *UserRepository) FindByUsernameAndPassword(username string, password string) (model.User, error) {
	var user model.User
	result := ur.Db.Where("username = ?", username).Where("password = ?", password).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, result.Error
}

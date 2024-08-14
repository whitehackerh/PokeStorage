package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

type (
	UserService struct {
		repo IUserRepository
	}
	IUserRepository interface {
		Create(model.User) error
	}
)

func NewUserService(userRepository IUserRepository) entity.IUserService {
	return &UserService{
		repo: userRepository,
	}
}

func (u *UserService) Create(user entity.User) (entity.User, error) {
	userModel := u.MapEntityToModel(user)
	err := u.repo.Create(userModel)
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *UserService) MapEntityToModel(user entity.User) model.User {
	return model.User{
		Id:           user.Id(),
		Username:     user.Username(),
		Password:     user.Password(),
		EmailAddress: user.EmailAddress(),
		Name:         user.Name(),
	}
}

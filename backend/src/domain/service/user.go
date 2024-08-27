package service

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

type (
	IUserService interface {
		Create(entity.User) error
		FindByUsernameAndPassword(string, string) (entity.User, error)
	}
	UserService struct {
		repo repository.IUserRepository
	}
)

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		repo: userRepository,
	}
}

func (u *UserService) Create(user entity.User) error {
	userModel := u.mapEntityToModel(user)
	return u.repo.Create(userModel)
}

func (u *UserService) mapEntityToModel(user entity.User) model.User {
	return model.User{
		Id:           user.Id(),
		Username:     user.Username(),
		Password:     user.Password(),
		EmailAddress: user.EmailAddress(),
		Name:         user.Name(),
	}
}

func (u *UserService) FindByUsernameAndPassword(username string, password string) (entity.User, error) {
	userModel, err := u.repo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return entity.User{}, err
	}
	return u.mapModelToEntity(userModel), err
}

func (u *UserService) mapModelToEntity(model model.User) entity.User {
	return entity.NewUser(
		model.Id,
		model.Username,
		model.Password,
		model.EmailAddress,
		model.Name,
	)
}

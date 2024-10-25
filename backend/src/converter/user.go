package converter

import (
	"github.com/whitehackerh/PokeStorage/src/domain/entity"
	"github.com/whitehackerh/PokeStorage/src/model"
)

func UserEntityToModel(user entity.User) model.User {
	return model.User{
		Id:           user.Id(),
		Username:     user.Username(),
		Password:     user.Password(),
		EmailAddress: user.EmailAddress(),
		Name:         user.Name(),
	}
}

func UserModelToEntity(user model.User) entity.User {
	return entity.NewUser(
		user.Id,
		user.Username,
		user.Password,
		user.EmailAddress,
		user.Name,
	)
}

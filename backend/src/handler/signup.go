package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func Signup(c *gin.Context) {
	var api = "/signup"
	var input usecase.SignupInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(api, err.Error()))
		return
	}
	uc := usecase.NewSignupInteractor(
		service.NewUserService(
			repository.NewUserRepository(
				infrastructure.ConnectDb(),
			),
		),
		middleware.NewAuth(),
		presenter.NewSignupPresenter(),
	)
	output, token, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(api, err.Error()))
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(200, presenter.NewCommonPresenter(api, output))
}

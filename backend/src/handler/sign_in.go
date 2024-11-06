package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func SignIn(c *gin.Context) {
	var url = "/sign-in"
	var input usecase.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}
	uc := usecase.NewSignInInteractor(
		repository.NewUserRepository(
			infrastructure.ConnectDb(),
		),
		middleware.NewAuth(repository.NewJwtBlacklistRepository(infrastructure.ConnectDb())),
		presenter.NewSignInPresenter(),
	)
	output, token, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

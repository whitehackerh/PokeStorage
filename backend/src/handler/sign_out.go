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

func SignOut(c *gin.Context) {
	var api = "/sign-out"
	var input = usecase.SignOutInput{Token: c.MustGet("token").(string), Claims: *c.MustGet("claims").(*middleware.Claims)}

	uc := usecase.NewSignOutInteractor(
		middleware.NewAuth(repository.NewJwtBlacklistRepository(infrastructure.ConnectDb())),
		presenter.NewSignOutPresenter(),
	)
	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(api, err.Error()))
		return
	}
	c.JSON(200, presenter.NewCommonPresenter(api, output))
}

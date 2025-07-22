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

func GetSwShTeams(c *gin.Context) {
	var url = "/swsh/teams"

	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	uc := usecase.NewGetSwShTeamsInteractor(
		service.NewSwShTeamService(
			service.NewTeamService(),
			service.NewSwShBredPokemonService(
				service.NewBredPokemonService(),
			),
		),
		repository.NewSwShTeamRepository(db),
		repository.NewSwShBredPokemonRepository(db),
		presenter.NewGetSwShTeamsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSwShTeamsInput{
		UserId: claims.UserId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

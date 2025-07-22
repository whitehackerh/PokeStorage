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

func GetSVTeams(c *gin.Context) {
	var url = "/sv/teams"

	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	uc := usecase.NewGetSVTeamsInteractor(
		service.NewSVTeamService(
			service.NewTeamService(),
			service.NewSVBredPokemonService(
				service.NewBredPokemonService(),
			),
		),
		repository.NewSVTeamRepository(db),
		repository.NewSVBredPokemonRepository(db),
		presenter.NewGetSVTeamsPresenter(),
	)

	output, err := uc.Execute(usecase.GetSVTeamsInput{
		UserId: claims.UserId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

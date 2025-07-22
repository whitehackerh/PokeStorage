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
	"gorm.io/gorm"
)

func PutSVTeams(c *gin.Context) {
	var url = "/sv/teams"

	claims := *c.MustGet("claims").(*middleware.Claims)

	db := infrastructure.ConnectDb()
	uc := usecase.NewPutSVTeamsInteractor(
		service.NewSVTeamService(
			service.NewTeamService(),
			service.NewSVBredPokemonService(
				service.NewBredPokemonService(),
			),
		),
		repository.NewSVTeamRepository(db),
		presenter.NewPutSVTeamsPresenter(),
	)

	input := usecase.PutSVTeamsInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, claims.UserId, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}

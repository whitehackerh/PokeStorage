package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
	"gorm.io/gorm"
)

func DeleteSVTeams(c *gin.Context) {
	var url = "/sv/teams"
	var input = usecase.DeleteSVTeamsInput{}

	input.TeamId = c.Param("id")

	db := infrastructure.ConnectDb()
	uc := usecase.NewDeleteSVTeamsInteractor(
		repository.NewSVTeamRepository(db),
		presenter.NewDeleteSVTeamsPresenter(),
	)

	infrastructure.WithTransaction(db, c, func(tx *gorm.DB) error {
		output, err := uc.Execute(input, tx)
		if err != nil {
			c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
			return err
		}

		c.JSON(201, presenter.NewCommonPresenter(url, output))
		return nil
	})
}

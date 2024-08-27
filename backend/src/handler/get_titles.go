package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/domain/service"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetTitles(c *gin.Context) {
	var api = "/titles"
	var input = usecase.GetTitlesInput{}

	uc := usecase.NewGetTitlesInteractor(
		service.NewTitleService(
			repository.NewTitleRepository(
				infrastructure.ConnectDb(),
			),
		),
		presenter.NewGetTitlesPresenter(),
	)
	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(api, err.Error()))
		return
	}
	c.JSON(200, presenter.NewCommonPresenter(api, output))
}

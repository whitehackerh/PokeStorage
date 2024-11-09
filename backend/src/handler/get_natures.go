package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetNatures(c *gin.Context) {
	var url = "/natures"
	var input = usecase.GetNaturesInput{}

	uc := usecase.NewGetNaturesInteractor(
		repository.NewNatureRepository(
			infrastructure.ConnectDb(),
		),
		presenter.NewGetNaturesPresenter(),
	)
	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

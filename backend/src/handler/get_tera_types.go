package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/presenter"
	"github.com/whitehackerh/PokeStorage/src/repository"
	"github.com/whitehackerh/PokeStorage/src/usecase"
)

func GetTeraTypes(c *gin.Context) {
	var url = "/tera-types"
	var input = usecase.GetTeraTypesInput{}

	uc := usecase.NewGetTeraTypesInteractor(
		repository.NewTeraTypeRepository(
			infrastructure.ConnectDb(),
		),
		presenter.NewGetTeraTypesPresenter(),
	)
	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

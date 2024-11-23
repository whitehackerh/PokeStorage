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

func GetSwShMoves(c *gin.Context) {
	var url = "/swsh/moves"
	var input = usecase.GetSwShMovesInput{}

	uc := usecase.NewSwShGetMovesInteractor(
		repository.NewSwShMoveRepository(
			infrastructure.ConnectDb(),
		),
		service.NewSwShMoveService(),
		presenter.NewGetSwShMovesPresenter(),
	)

	output, err := uc.Execute(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.NewCommonPresenter(url, err.Error()))
		return
	}

	c.JSON(200, presenter.NewCommonPresenter(url, output))
}

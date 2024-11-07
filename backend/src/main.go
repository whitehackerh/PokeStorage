package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/whitehackerh/PokeStorage/src/handler"
	"github.com/whitehackerh/PokeStorage/src/infrastructure"
	"github.com/whitehackerh/PokeStorage/src/middleware"
	"github.com/whitehackerh/PokeStorage/src/repository"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.POST("/sign-up", handler.SignUp)
		api.POST("/sign-in", handler.SignIn)

		auth := middleware.NewAuth(repository.NewJwtBlacklistRepository(infrastructure.ConnectDb()))
		authGroup := api.Group("/auth")
		authGroup.Use(auth.ParseToken())
		{
			authGroup.POST("/sign-out", handler.SignOut)
			authGroup.GET("/titles", handler.GetTitles)
			authGroup.GET("/pokemons/:title-id", handler.GetPokemons)
			authGroup.GET("/tera-types", handler.GetTeraTypes)
			// /items
			// /moves
		}
	}
	r.Run()
}

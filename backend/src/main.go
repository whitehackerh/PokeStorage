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

			authGroup.GET("/tera-types", handler.GetTeraTypes)
			authGroup.GET("/natures", handler.GetNatures)
			authGroup.GET("/genders", handler.GetGenders)

			swshGroup := authGroup.Group("/swsh")
			{
				swshGroup.GET("/pokemons", handler.GetSwShPokemons)
				swshGroup.GET("/items", handler.GetSwShItems)
				swshGroup.GET("/moves", handler.GetSwShMoves)
				swshGroup.POST("/bred-pokemons", handler.PostSwShBredPokemons)
				swshGroup.GET("/bred-pokemons", handler.GetSwShBredPokemons)
			}

			svGroup := authGroup.Group("/sv")
			{
				svGroup.GET("/pokemons", handler.GetSVPokemons)
				svGroup.GET("/items", handler.GetSVItems)
				svGroup.GET("/moves", handler.GetSVMoves)
				svGroup.POST("/bred-pokemons", handler.PostSVBredPokemons)
				svGroup.GET("/bred-pokemons", handler.GetSVBredPokemons)
			}
		}
	}
	r.Run()
}

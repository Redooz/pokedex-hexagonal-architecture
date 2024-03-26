package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redooz/podekex-hexagonal-architecture/infrastructure"
)

func main() {
	router := gin.Default()

	pokedexController := infrastructure.InitializePokedexController()
	pokedexController.InitRoutes(router)

	typeController := infrastructure.InitializeTypeController()
	typeController.InitRoutes(router)

	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}

}

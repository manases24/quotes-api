package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mnsh5/quotes/controller"
	"github.com/mnsh5/quotes/database"
	"github.com/mnsh5/quotes/services"
)

var router = gin.Default()

// Run iniciará el servidor
func Run() {
	setupRoutes()
	router.Run(":2024")
}

func setupRoutes() {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db := database.DB
	quotesService := services.NewQuotesServiceImpl(db)
	quote := controller.NewQuotesController(quotesService)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/quote", quote.Create)
	}
}

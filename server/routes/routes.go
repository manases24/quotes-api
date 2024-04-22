package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mnsh5/quotes/controller"
	"github.com/mnsh5/quotes/database"
	"github.com/mnsh5/quotes/services"
)

func SetupRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db := database.DB
	quotesService := services.NewQuotesServiceImpl(db)
	quote := controller.NewQuotesController(quotesService)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.POST("/quote", quote.Create)

	return router
}

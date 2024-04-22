package services

import "github.com/gin-gonic/gin"

type QuotesService interface {
	FindAll(c *gin.Context)
	FindRandomQuote(c *gin.Context)
	Create(c *gin.Context)
}

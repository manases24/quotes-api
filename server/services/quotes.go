package services

import "github.com/gin-gonic/gin"

type QuotesService interface {
	Create(c *gin.Context) error
}

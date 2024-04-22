package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mnsh5/quotes/services"
)

type QuotesController struct {
	quotesService services.QuotesService
}

func NewQuotesController(s services.QuotesService) *QuotesController {
	return &QuotesController{quotesService: s}
}

func (q QuotesController) Create(c *gin.Context) {
	q.quotesService.Create(c)
}

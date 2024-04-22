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

func (q QuotesController) FindRandomQuote(c *gin.Context) {
	q.quotesService.FindRandomQuote(c)
}

func (q QuotesController) FindAll(c *gin.Context) {
	q.quotesService.FindAll(c)
}

func (q QuotesController) FindById(c *gin.Context) {
	q.quotesService.FindById(c)
}

func (q QuotesController) Create(c *gin.Context) {
	q.quotesService.Create(c)
}

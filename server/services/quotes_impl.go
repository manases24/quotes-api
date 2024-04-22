package services

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/mnsh5/quotes/models"
	"gorm.io/gorm"
)

type QuotesServiceImpl struct {
	Db *gorm.DB
}

func NewQuotesServiceImpl(Db *gorm.DB) QuotesService {
	return &QuotesServiceImpl{Db: Db}
}

func (q *QuotesServiceImpl) Create(c *gin.Context) {
	log.Println("start to execute program add data quote")
	var quote models.Quotes

	// Bind the request body to the quote struct
	if err := c.Bind(&quote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to parse request body": err.Error()})
		c.Error(err)
		return
	}

	// Create the quote in the database
	if err := q.Db.Create(&quote).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Failed to create task": err.Error()})
		c.Error(err)
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Quote has been created",
		"data":    quote,
	})
}

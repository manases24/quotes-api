package services

import (
	"math/rand"
	"net/http"
	"time"

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

func (q *QuotesServiceImpl) FindAll(c *gin.Context) {
	var quotes []models.Quotes
	if err := q.Db.Find(&quotes).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No quotes found"})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "All quotes retrieved successfully",
		"quote":   quotes,
	})
}

func (q *QuotesServiceImpl) FindRandomQuote(c *gin.Context) {
	// Inicializar el generador de números aleatorios
	// Crear un generador de números aleatorios local
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	// Contar el total de quotes en la base de datos
	var count int64
	if err := q.Db.Model(&models.Quotes{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Verificar si no hay quotes en la base de datos
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No quotes found"})
		return
	}

	// Generar un ID aleatorio dentro del rango de IDs disponibles
	randID := rnd.Int63n(count) + 1 // Sumamos 1 para evitar el ID 0 si comienza desde 1

	// Buscar el quote correspondiente al ID aleatorio
	var randomQuote models.Quotes
	if err := q.Db.First(&randomQuote, randID).Error; err != nil {
		// Si el ID aleatorio no corresponde a un quote existente, intentar otra vez
		if err == gorm.ErrRecordNotFound {
			q.FindRandomQuote(c)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver el quote aleatorio como respuesta JSON
	c.JSON(http.StatusOK, gin.H{
		"quote":  randomQuote.Quote,
		"author": randomQuote.Author,
	})
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
		"quote":   quote,
	})
}

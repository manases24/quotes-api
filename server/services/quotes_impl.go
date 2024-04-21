package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuotesServiceImpl struct {
	Db *gorm.DB
}

func NewQuotesServiceImpl(Db *gorm.DB) QuotesService {
	return &QuotesService{Db: Db}
}

func (q *QuotesServiceImpl) Create(c *gin.Context) error {

}

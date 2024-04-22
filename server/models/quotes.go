package models

import "gorm.io/gorm"

type Quotes struct {
	gorm.Model
	Quote string `gorm:"not null" validate:"required,min=3,max=200" json:"quote"`
}

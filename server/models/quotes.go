package models

import "gorm.io/gorm"

type Quotes struct {
	gorm.Model
	Author string `gorm:"not null" validate:"required,min=3,max=10" json:"author"`
	Quote  string `gorm:"not null" validate:"required,min=3,max=200" json:"quote"`
}

package models

import (
	"github.com/jinzhu/gorm"
)

// Book struct
type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// TableName create the table in db
func (b *Book) TableName() string {
	return "book"
}

package models

import (
	_ "github.com/jinzhu/gorm"
)

//Book model structure
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

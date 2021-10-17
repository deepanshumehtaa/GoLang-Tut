package models

import (
	"github.com/deepanshumehtaa/bookstore-go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Auto migration Funciton for the db model
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Logic
//

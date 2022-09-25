package models

import (
	"github.com/bopepsi/p3-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

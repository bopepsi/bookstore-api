package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=bopepsi dbname=books password=")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("connected db")
	return d
}

func GetDB() *gorm.DB {
	return db
}

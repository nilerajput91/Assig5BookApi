package main

import (
	"github.com/nilerajput91/assig5bookapi/config"
	"github.com/nilerajput91/assig5bookapi/models"
	"github.com/nilerajput91/assig5bookapi/routers"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {
	config.DB, err = gorm.Open("postgres", "postgres://@localhost/")

	if err != nil {
		fmt.Println("status:", err)

	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Book{})

	r := routers.SetupRouter()

	r.Run()
}

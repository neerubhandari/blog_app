package main

import (
	"fmt"

	"blog_app/Config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// store the database connection in the Config.DB variable for use throughout the program
	Config.DB = db

	// Create a new user

	fmt.Println("Connected to database")
}

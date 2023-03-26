package main

import (
	"fmt"

	"blog_app/Config"
	controllers "blog_app/Controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	db, err := gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// store the database connection in the Config.DB variable for use throughout the program
	// Config.DB = db

	// create a new PostController instance
	postController := controllers.PostController{DB: db}
	fmt.Println(postController)

	fmt.Println("Connected to database")
	// map the GetAllPosts function to a GET request for the "/posts" endpoint
	router.GET("/posts", postController.GetAllPosts)
	router.Run(":8081")

}

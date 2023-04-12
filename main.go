package main

import (
	"blog_app/Config"
	controllers "blog_app/Controllers"
	models "blog_app/Model"
	"fmt"

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
	Config.DB = db

	// create a new PostController instance
	postController := controllers.PostController{DB: db}
	fmt.Println(postController)

	fmt.Println("Connected to database")

	// Migrate the Post model to the database
	db.AutoMigrate(&models.Post{})

	// Inject the database connection into the Gin context using middleware
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	// map the GetAllPosts function to a GET request for the "/posts" endpoint
	router.GET("/posts", postController.GetAllPosts)

	// Register the CreatePost function as a handler
	router.POST("/posts", controllers.CreatePost)

	//GET PARTICULAR POST BASED ON THE ID
	router.GET("/posts/:id", controllers.FindPost)

	//DELETE THE POST
	router.DELETE("/posts/:id", controllers.DeletePost) // new
	router.Run(":8084")

}

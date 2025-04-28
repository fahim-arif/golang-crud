package main

import (
	"main/controllers"
	"main/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostCreate)

	router.GET("/posts", controllers.PostsIndex)

	router.GET("/posts/:id", controllers.PostsShow)

	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run()
}

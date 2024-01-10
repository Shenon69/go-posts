package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
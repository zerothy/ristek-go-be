package main

/*
CompileDaemon -command="./server"
*/

import (
	"server/controllers"
	"server/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Your existing routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yey Jalannn",
		})
	})

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	// r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts", controllers.PostsIndexByType)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}

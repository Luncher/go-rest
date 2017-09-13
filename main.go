package main

import (
	"github.com/Luncher/go-rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		movie := new(controllers.UserController)
		v1.POST("/movies", movie.Create)
		v1.GET("/movies/:id", movie.Get)
		v1.PUT("/movies/:id", movie.Update)
		v1.DELETE("/movies/:id", movie.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.NotFound, "Not Found")
	})

	router.Run()
}

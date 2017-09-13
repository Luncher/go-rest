package main

import (
	"bytes"
	"encoding/json"
	"github.com/Luncher/go-rest/controllers"
	"github.com/Luncher/go-rest/forms"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Movie struct {
	Name   string
	Desc   string
	Rating float32
}

func SetupRouter() *gin.Engine {
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
		c.String(http.StatusNotFound, "Not Found")
	})

	return router
}

func TestCreateMovie(t *testing.T) {
	testRouter := SetupRouter()

	movie := forms.CreateMovieCommand{}
	movie.Name = "foobar"
	movie.Desc = "the foobar movie"
	movie.Rating = 6

	body, _ := json.Marshal(movie)
	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", "/v1/movies", buf)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatal("Invalid Request")
	}
}

func TestGetMovie(t *testing.T) {
	testRouter := SetupRouter()

}

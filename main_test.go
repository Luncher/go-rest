package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Luncher/go-rest/controllers"
	"github.com/Luncher/go-rest/forms"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Movie struct {
	id     string
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
		v1.GET("/movies", movie.Find)
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
	fmt.Println(resp.Body.String())
	if resp.Code != 200 {
		t.Fatal("Invalid Request")
	}
}

var movieId string

func TestFindAllMovie(t *testing.T) {
	testRouter := SetupRouter()
	req, err := http.NewRequest("GET", "/v1/movies", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	result := struct {
		Data []struct {
			Id string `json:"id"`
		}
	}{}
	json.Unmarshal(buf, &result)
	movieId = result.Data[0].Id
	if resp.Code != 200 {
		t.Fatal("Invalid Request", resp.Code, resp.Body.String())
	}
}

func TestUpdateMovie(t *testing.T) {
	testRouter := SetupRouter()
	data := forms.UpdateMovieCommand{Name: "titanic", Rating: 10, Desc: "hello"}
	body, _ := json.Marshal(data)
	buf := bytes.NewBuffer(body)
	uri := fmt.Sprintf("/v1/movies/%s", "59b8e1f73d27fe6828a82bf3")
	fmt.Println(uri)
	req, err := http.NewRequest("PUT", uri, buf)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatal("Invalid Request", resp.Code, resp.Body.String())
	}
}

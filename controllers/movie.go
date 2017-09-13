package controllers

import (
	"github.com/Luncher/go-rest/forms"
	"github.com/Luncher/go-rest/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

var movieModel = new(models.MovieModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data forms.CreateMovieCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	err := movieModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Movie created"})
}

func (user *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := movieModel.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Movie not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

func (user *UserController) Find(c *gin.Context) {
	list, err := movieModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (user *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	data := forms.UpdateMovieCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}

	err := movieModel.Update(id, data)
	if err != nil {
		c.JSON(406, gin.H{"message": "movie count not be updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Movie updated"})
}

func (user *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := movieModel.Delete(id)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie could not be deleted", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Movie deleted"})
}

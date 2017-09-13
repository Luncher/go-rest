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

	movieId, err := movieModel.Ceate(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Movie could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Movie created", "id": movieId})
}

func (user *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(404, gin.H{"message": "Invalid parameter"})
	} else {
		profile, err := movieModel.findOne(id)
		if err != nil {
			c.JSON(404, gin.H{"message": "Movie not found", "error": err.Error()})
			c.Abort()
		} else {
			c.JSON(200, gin.H{"data": profile})
		}
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

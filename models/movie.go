package models

import (
	"errors"
	"github.com/Luncher/go-rest/db"
	"github.com/Luncher/go-rest/forms"
	"time"
)

type Movie struct {
	Name   string
	Rating float32
	Desc   string
}

type MovieModel struct{}

var dbConnect = db.NewConnection("localhost")

func (m MovieModel) Create(data forms.CreateMovieCommand) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.Insert(&Movie{
		Name: data.Name,
		Desc: data.Desc,
		Rating: data.Rating
	})

	return err
}

func (m MovieModel) Get(id string) (movie Movie, err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.FindId(id).One(&movie)

	return movie, err
}

func (m MovieModel) Update(id string, data forms.UpdateMovieCommand) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.UpdateId(id, data)

	return err
}

func (m MovieModel) Delete(id string) (err error) {
	collection := dbConnect.Use("test-mgo", "movies")
	err = collection.RemoveId(id)	

	return err
}

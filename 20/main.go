package main

import (
	"github.com/kirigaikabuto/golangLessons1900/20/movie"
	"log"
)

func main() {
	movieMongoStore, err := movie.NewMovieStore(movie.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "movies",
	})
	if err != nil {
		log.Fatal(err)
	}
}

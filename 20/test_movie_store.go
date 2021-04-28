package main

import (
	"fmt"
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
	m1 := &movie.Movie{
		Name:        "movie2",
		Description: "movie2",
		ImageUrl:    "movie2",
	}
	m1, err = movieMongoStore.Create(m1)
	if err != nil {
		log.Fatal(err)
	}
	movies, err := movieMongoStore.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movies)
	//m2, err := movieMongoStore.Get(34)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(m2)
	//err = movieMongoStore.Delete(34)
	//m2, err = movieMongoStore.Get(34)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(m2)
}

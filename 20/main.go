package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/20/movie"
	"log"
	"net/http"
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
	httpEndpoints := movie.NewHttpEndpoints(movieMongoStore)
	http.HandleFunc("/", httpEndpoints.MainPage)
	http.HandleFunc("/add_page", httpEndpoints.AddPage)
	http.HandleFunc("/add_page_action", httpEndpoints.AddPageAction)
	http.HandleFunc("/detail_page", httpEndpoints.DetailPage)
	fmt.Println("start server on port 8080")
	http.ListenAndServe(":8080", nil)
}

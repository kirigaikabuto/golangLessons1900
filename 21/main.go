package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/golangLessons1900/21/movie"
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
	router := mux.NewRouter()
	httpEndpoints := movie.NewHttpEndpoints(movieMongoStore)
	router.Methods("GET").Path("/").HandlerFunc(httpEndpoints.MainPage())
	router.Methods("GET").Path("/add_page").HandlerFunc(httpEndpoints.AddPage())
	router.Methods("POST").Path("/add_page_action").HandlerFunc(httpEndpoints.AddPageAction())
	router.Methods("GET").Path("/detail_page/{id}").HandlerFunc(httpEndpoints.DetailPage("id"))
	router.Methods("GET").Path("/delete/{idParam}").HandlerFunc(httpEndpoints.DeleteAction("idParam"))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/golangLessons1900/23/config"
	"github.com/kirigaikabuto/golangLessons1900/23/start"
	"github.com/kirigaikabuto/golangLessons1900/23/users"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	usersMongoStore, err := users.NewUsersStore(config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "users",
	})
	if err != nil {
		log.Fatal(err)
	}
	mainEndpoints := start.NewHttpEndpoints(usersMongoStore)
	router.Methods("GET").Path("/test").HandlerFunc(mainEndpoints.TestEndpoint())
	router.Methods("GET").Path("/test/{param}").HandlerFunc(mainEndpoints.TestEndpointWithParam("param"))
	router.Methods("POST").Path("/test").HandlerFunc(mainEndpoints.TestPostEndpoint())
	router.Methods("POST").Path("/register").HandlerFunc(mainEndpoints.RegisterEndpoint())
	router.Methods("POST").Path("/login").HandlerFunc(mainEndpoints.LoginEndpoint())
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/golangLessons1900/23/start"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	mainEndpoints := start.NewHttpEndpoints()
	router.Methods("GET").Path("/test").HandlerFunc(mainEndpoints.TestEndpoint())
	router.Methods("GET").Path("/test/{param}").HandlerFunc(mainEndpoints.TestEndpointWithParam("param"))
	router.Methods("POST").Path("/test").HandlerFunc(mainEndpoints.TestPostEndpoint())
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

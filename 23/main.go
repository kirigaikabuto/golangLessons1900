package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/golangLessons1900/23/common"
	"github.com/kirigaikabuto/golangLessons1900/23/config"
	"github.com/kirigaikabuto/golangLessons1900/23/products"
	"github.com/kirigaikabuto/golangLessons1900/23/redis_lib"
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
	redisStore, err := redis_lib.NewRedisConnect(config.RedisConfig{
		Host: "localhost",
		Port: "6379",
	})
	if err != nil {
		log.Fatal(err)
	}

	//middleware connect
	middleware := common.NewMiddleware(redisStore)

	mainEndpoints := start.NewHttpEndpoints(usersMongoStore, redisStore)
	//start
	//router.Methods("GET").Path("/test").HandlerFunc(mainEndpoints.TestEndpoint())
	//router.Methods("GET").Path("/test/{param}").HandlerFunc(mainEndpoints.TestEndpointWithParam("param"))
	//router.Methods("POST").Path("/test").HandlerFunc(mainEndpoints.TestPostEndpoint())
	router.Methods("POST").Path("/register").HandlerFunc(mainEndpoints.RegisterEndpoint())
	router.Methods("POST").Path("/login").HandlerFunc(mainEndpoints.LoginEndpoint())
	router.Methods("GET").Path("/profile").Handler(middleware.LoginMiddleware(http.HandlerFunc(mainEndpoints.ProfileEndpoint())))

	//products
	productMongoStore, err := products.NewProductStore(config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "products",
	})
	productHttpEndpoints := products.NewProductHttpEndpoints(productMongoStore)

	router.Methods("POST").Path("/products").Handler(middleware.LoginMiddleware(http.HandlerFunc(productHttpEndpoints.CreateProduct())))
	router.Methods("GET").Path("/products").Handler(middleware.LoginMiddleware(http.HandlerFunc(productHttpEndpoints.ListProduct())))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

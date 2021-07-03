package main

import (
	"fmt"
	"github.com/gorilla/mux"
	common "github.com/kirigaikabuto/common-lib31"
	orders "github.com/kirigaikabuto/orders31"
	products "github.com/kirigaikabuto/products31"
	redis_lib "github.com/kirigaikabuto/common-lib31"
	start "github.com/kirigaikabuto/start31"
	users "github.com/kirigaikabuto/users31"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	usersMongoStore, err := users.NewUsersStore(common.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "users",
	})
	if err != nil {
		log.Fatal(err)
	}
	redisStore, err := redis_lib.NewRedisConnect(common.RedisConfig{
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
	productMongoStore, err := products.NewProductStore(common.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "products",
	})
	if err != nil {
		log.Fatal(err)
	}
	productHttpEndpoints := products.NewProductHttpEndpoints(productMongoStore)

	router.Methods("POST").Path("/products").Handler(middleware.LoginMiddleware(http.HandlerFunc(productHttpEndpoints.CreateProduct())))
	router.Methods("GET").Path("/products").Handler(middleware.LoginMiddleware(http.HandlerFunc(productHttpEndpoints.ListProduct())))

	ordersMongoStore, err := orders.NewOrdersStore(common.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "orders",
	})
	if err != nil {
		log.Fatal(err)
	}
	ordersHttpEndpoints := orders.NewOrderHttpEndpoints(ordersMongoStore)

	router.Methods("POST").Path("/orders").Handler(middleware.LoginMiddleware(http.HandlerFunc(ordersHttpEndpoints.CreateOrder())))
	router.Methods("GET").Path("/orders").Handler(middleware.LoginMiddleware(http.HandlerFunc(ordersHttpEndpoints.ListOrder())))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

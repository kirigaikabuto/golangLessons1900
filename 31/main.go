package main

import (
	"fmt"
	"github.com/djumanoff/amqp"
	"github.com/gorilla/mux"
	common "github.com/kirigaikabuto/common-lib31"
	redis_lib "github.com/kirigaikabuto/common-lib31"
	orders "github.com/kirigaikabuto/orders31"
	products "github.com/kirigaikabuto/products31"
	start "github.com/kirigaikabuto/start31"
	"log"
	"net/http"
)

func main() {


	rabbitConfig := amqp.Config{
		Host:     "localhost",
		Port:     5672,
		LogLevel: 5,
	}
	sess := amqp.NewSession(rabbitConfig)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}


	redisStore, err := redis_lib.NewRedisConnectStore(common.RedisConfig{
		Host: "localhost",
		Port: "6379",
	})
	if err != nil {
		log.Fatal(err)
	}

	//middleware connect
	middleware := common.NewMiddleware(*redisStore)

	mainEndpoints := start.NewHttpEndpoints(clt, redisStore)
	//start

	router := mux.NewRouter()

	router.Methods("POST").Path("/register").HandlerFunc(mainEndpoints.RegisterEndpoint())
	router.Methods("POST").Path("/login").HandlerFunc(mainEndpoints.LoginEndpoint())
	router.Methods("GET").Path("/profile").HandlerFunc(middleware.LoginMiddleware(mainEndpoints.ProfileEndpoint()))

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

	router.Methods("POST").Path("/products").HandlerFunc(middleware.LoginMiddleware(productHttpEndpoints.CreateProduct()))
	router.Methods("GET").Path("/products").HandlerFunc(middleware.LoginMiddleware(productHttpEndpoints.ListProduct()))

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

	router.Methods("POST").Path("/orders").HandlerFunc(middleware.LoginMiddleware(ordersHttpEndpoints.CreateOrder()))
	router.Methods("GET").Path("/orders").HandlerFunc(middleware.LoginMiddleware(ordersHttpEndpoints.ListOrder()))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

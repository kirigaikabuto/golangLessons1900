package main

import (
	"fmt"
	"github.com/djumanoff/amqp"
	"github.com/gorilla/mux"
	common "github.com/kirigaikabuto/common-lib31"
	redis_lib "github.com/kirigaikabuto/common-lib31"
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
	router.Methods("POST").Path("/product/create").HandlerFunc(middleware.LoginMiddleware(mainEndpoints.CreateProductEndpoint()))
	router.Methods("GET").Path("/product/list").HandlerFunc(middleware.LoginMiddleware(mainEndpoints.ListProductEndpoint()))


	router.Methods("POST").Path("/orders/create").HandlerFunc(middleware.LoginMiddleware(mainEndpoints.CreateOrder()))
	router.Methods("GET").Path("/orders/list").HandlerFunc(middleware.LoginMiddleware(mainEndpoints.ListOrder()))
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

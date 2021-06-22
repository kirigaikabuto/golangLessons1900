package main

import (
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/golangLessons1900/29/products"
)

func main() {
	//open access to function getProducts

	//config for rabbitmq
	config := amqp.Config{
		Host:        "localhost",
		VirtualHost: "",
		Port:        5672,
		User:        "",
		Password:    "",
		LogLevel:    5,
	}
	serverConfig := amqp.ServerConfig{
		ResponseX: "response",
		RequestX:  "request",
	}
	//connect to rabbitmq
	sess := amqp.NewSession(config)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	//create server or api for internal request(GRPC)
	srv, err := sess.Server(serverConfig)
	if err != nil {
		panic(err)
	}

	amqpProductEndpoints := products.NewAmqpEndpoints()
	srv.Endpoint("amqp_get_product", amqpProductEndpoints.GetProductAmqpEndpoint())

	if err := srv.Start(); err != nil {
		panic(err)
		return
	}

}

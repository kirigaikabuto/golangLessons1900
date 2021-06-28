package main

import (
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/golangLessons1900/30/users"
)

func main() {
	usersMongoStore, err := users.NewUsersStore(users.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "users",
	})
	if err != nil {
		panic(err)
		return
	}
	usersAmqpEndpoints := users.NewUsersAmqpEndpoints(usersMongoStore)
	rabbitConfig := amqp.Config{
		Host:     "localhost",
		Port:     5672,
		LogLevel: 5,
	}
	serverConfig := amqp.ServerConfig{
		ResponseX: "response",
		RequestX:  "request",
	}

	sess := amqp.NewSession(rabbitConfig)
	err = sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	srv, err := sess.Server(serverConfig)
	if err != nil {
		panic(err)
		return
	}
	srv.Endpoint("list_users", usersAmqpEndpoints.ListUsers())
	err = srv.Start()
	if err != nil {
		panic(err)
		return
	}

}

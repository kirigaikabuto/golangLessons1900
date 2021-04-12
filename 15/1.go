package main

import (
	"context"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/15/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//go get go.mongodb.org/mongo-driver/mongo
func main() {
	//connect to database
	//create config
	cf := config.MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		Database: "testdb",
	}
	//config for client
	clientOptions := options.Client().ApplyURI("mongodb://" + cf.Host + ":" + cf.Port)
	//connect to database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//check for connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//connect to database
	db := client.Database(cf.Database)
	//connect to table/collection
	collection := db.Collection("users")
	fmt.Println(collection)
}

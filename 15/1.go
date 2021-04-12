package main

import (
	"context"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/15/config"
	"github.com/kirigaikabuto/golangLessons1900/15/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

//go get go.mongodb.org/mongo-driver/mongo
func main() {
	//connect to database
	//create config
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
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
	//create collection
	err = db.CreateCollection(context.TODO(), cf.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		log.Fatal(err.Error())
	}
	//create connection to collection
	collection := db.Collection(cf.CollectionName)
	//insert first element
	user1 := models.User{
		Username: "user2",
		Password: "1232131",
	}
	result, err := collection.InsertOne(context.TODO(), user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	//get all data from db
	users := []models.User{}
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range users {
		fmt.Println(v.Username, v.Password)
	}
}

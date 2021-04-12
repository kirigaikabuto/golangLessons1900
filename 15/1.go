package main

import "github.com/kirigaikabuto/golangLessons1900/15/config"

//go get go.mongodb.org/mongo-driver/mongo
func main() {
	//connect to database
	//create config
	cf := config.MongoConfig{
		Host:     "localhost",
		Port:     "27017",
		Database: "test_db_12",
	}
}

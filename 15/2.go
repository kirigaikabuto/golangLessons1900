package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/15/config"
	"github.com/kirigaikabuto/golangLessons1900/15/models"
	"log"
)

func main() {
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	collection, err := models.GetConnection(cf)
	if err != nil {
		log.Fatal(err)
	}
	//user1 := models.User{
	//	Username: "user3",
	//	Password: "user3",
	//}
	//err = AddUser(collection, user1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	users, err := models.GetAllUsers(collection)
	fmt.Println(users)
}

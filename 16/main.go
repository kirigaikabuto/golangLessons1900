package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/16/config"
	models "github.com/kirigaikabuto/golangLessons1900/16/users"
	"log"
)

func main() {
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	usersStore, err := models.NewUsersStore(cf)
	if err != nil {
		log.Fatal(err)
	}
	user1 := models.User{
		Username: "123231",
		Password: "12323213",
	}
	err = usersStore.AddUser(user1)
	if err != nil {
		log.Fatal(err)
	}
	users, err := usersStore.List()
	for _, v := range users {
		fmt.Println(v)
	}
}

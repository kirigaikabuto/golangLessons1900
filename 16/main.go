package main

import (
	"errors"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/16/config"
	"github.com/kirigaikabuto/golangLessons1900/16/users"
	"log"
)

func main() {
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	username := "123231"
	password := "12323213"
	user := users.User{
		Username: username,
		Password: password,
	}
	usersStore, err := users.NewUsersStore(cf)
	if err != nil {
		log.Fatal(err)
	}
	usersData, err := usersStore.List()
	isExist := false
	for _, v := range usersData {
		if v.Username == user.Username && v.Password == v.Password {
			user = v
			isExist = true
		}
	}
	if isExist {
		fmt.Printf("Welcome %d", user.Id)
	} else {
		log.Fatal(errors.New("no user by this username and password"))
	}
}

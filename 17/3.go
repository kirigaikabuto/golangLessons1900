package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/17/config"
	"github.com/kirigaikabuto/golangLessons1900/17/users"
	"log"
)

func main() {
	user1 := users.User{
		Id:       123,
		Username: "asdsadsa",
		Password: "1232131",
	}
	//mongodb
	cf := config.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "testdb",
		CollectionName: "users",
	}
	usersStore, err := users.NewUsersStore(cf)
	if err != nil {
		log.Fatal(err)
	}
	err = usersStore.AddUser(user1)
	if err != nil {
		log.Fatal(err)
	}
	mongoUser, err := usersStore.GetById(56)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mongoUser)

	//postgre
	postgreConfig := config.Config{
		Host:             "localhost",
		User:             "kirito",
		Password:         "passanya99!",
		Database:         "golangExample",
		Port:             5432,
		ConnectionString: "",
		Params:           "sslmode=disable",
	}
	postgreUsersStore, err := users.NewUsersPostgreStore(postgreConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = postgreUsersStore.AddUser(user1)
	if err != nil {
		log.Fatal(err)
	}
	postgreUser, err := postgreUsersStore.GetById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(postgreUser)
	//postgresl
	//mysql
	//oralceplsql
}

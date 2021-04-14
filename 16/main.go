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
		Id:       661,
		Username: "1111111111111",
	}
	err = usersStore.Update(user1)
	if err != nil {
		log.Fatal(err)
	}
	oneUser, err := usersStore.GetById(661)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneUser)
	err = usersStore.Delete(661)
	if err != nil {
		log.Fatal(err)
	}
	oneUser, err = usersStore.GetById(661)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneUser)
	err = usersStore.AddUser(user1)
	if err != nil {
		log.Fatal(err)
	}
	oneUser, err = usersStore.GetById(661)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oneUser)
	users, err := usersStore.List()
	for _, v := range users {
		fmt.Println(v)
	}
}

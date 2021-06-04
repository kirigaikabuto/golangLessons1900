package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/config27"
	"github.com/kirigaikabuto/users27"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	store, err := users27.NewUsersStore(config27.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "ivi",
		CollectionName: "users",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	//user, err := store.GetByUsernameAndPassword("kirito", "i77asdsda")
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(user)
	users, err := store.List()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)
	fmt.Println("server is running on port 8080:")
	http.ListenAndServe(":8080", router)
}

package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/19/login"
	"log"
	"net/http"
)

func main() {
	usersPostgre, err := login.NewUsersStore(login.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "user_store",
		CollectionName: "new_users",
	})
	if err != nil {
		log.Fatal(err)
	}
	loginHttpEndpoints := login.NewLoginHttpEndpoints(usersPostgre)
	http.HandleFunc("/", loginHttpEndpoints.MainPage)
	http.HandleFunc("/template", loginHttpEndpoints.TemplateExample)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

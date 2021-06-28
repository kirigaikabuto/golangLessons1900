package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func main() {
	rabbitConfig := amqp.Config{
		Host:     "localhost",
		Port:     5672,
		LogLevel: 5,
	}
	sess := amqp.NewSession(rabbitConfig)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}
	userForCreate := &User{
		Username:  "123",
		Password:  "123",
		FirstName: "123",
		LastName:  "123",
		Avatar:    "123",
	}
	createUserJson, err := json.Marshal(userForCreate)
	if err != nil {
		panic(err)
		return
	}
	_, err = clt.Call("create_user", amqp.Message{Body: createUserJson})
	if err != nil {
		panic(err)
		return
	}
	users := []User{}
	response, err := clt.Call("list_users", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(response.Body, &users)
	if err != nil {
		panic(err)
		return
	}
	for _, v := range users {
		fmt.Println(v.Id, v.Username)
	}
}

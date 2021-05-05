package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type Student1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	key := "st1"
	val, err := client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	st := &Student1{}
	err = json.Unmarshal([]byte(val), &st)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(st.Username, st.Password)
}

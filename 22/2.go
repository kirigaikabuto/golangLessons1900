package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Student struct {
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
	value := &Student{
		Username: "kirito",
		Password: "1998",
	}
	dataJson, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Set(key, dataJson, time.Minute*1).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}

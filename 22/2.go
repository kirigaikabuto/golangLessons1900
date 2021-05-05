package main

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	key := "name"
	value := "yerassyl"
	err = client.Set(key, value, time.Minute*1).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}

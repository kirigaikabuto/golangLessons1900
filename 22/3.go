package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
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
	val, err := client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}

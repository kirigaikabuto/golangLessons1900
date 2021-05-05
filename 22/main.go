package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/22/redis_lib"
	"log"
	"time"
)

type Book struct {
	Name  string
	Price int
}

func main() {
	config := redis_lib.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	redisConnect, err := redis_lib.NewRedisConnect(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	book1 := &Book{
		Name:  "book1",
		Price: 130,
	}
	key := "books"
	err = redisConnect.SetValue(key, book1, time.Minute*1)
	if err != nil {
		log.Fatal(err)
		return
	}
	book2 := &Book{}
	err = redisConnect.GetValue(key, &book2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(book2)
}

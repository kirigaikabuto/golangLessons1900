package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "http://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	t1 := &Todo{}
	err = json.Unmarshal(data, &t1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1.Id, t1.Title, t1.UserId, t1.Completed)
	fmt.Println(t1)
}

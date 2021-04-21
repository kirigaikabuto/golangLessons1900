package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo3 struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "http://jsonplaceholder.typicode.com/todos"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	todos := []Todo3{}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range todos {
		fmt.Println(v)
	}
}

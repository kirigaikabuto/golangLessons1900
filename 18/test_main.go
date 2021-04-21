package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Student3 struct {
	Name string `json:"name"`
}

func main() {
	url := "http://127.0.0.1:8080/students"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	students := []Student3{}
	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range students {
		fmt.Println(v.Name)
	}
}

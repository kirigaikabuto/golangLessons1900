package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	post1 := &Post{
		UserId: 123,
		Title:  "asdasds",
		Body:   "adsasd",
	}
	jsonData, err := json.Marshal(post1)
	if err != nil {
		log.Fatal(err)
	}
	postData := bytes.NewReader(jsonData)
	response, err := http.Post(url, "application/json; charset=UTF-8", postData)
	if err != nil {
		log.Fatal(err)
	}
	responsePost := &Post{}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(responseData, &responsePost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(responsePost.Id)
}

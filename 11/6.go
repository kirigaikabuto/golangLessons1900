package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/11/models"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	var students []models.Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		panic(err)
	}
	for _, v := range students {
		fmt.Println(v.FirstName)
	}
}

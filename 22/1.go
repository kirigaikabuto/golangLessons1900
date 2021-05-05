package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	d := map[string]int{
		"name": 1,
		"age":  22,
	}
	d["surname"] = 123
	d["surname"] = 456
	println(d["age"])
	println(d["surname"])
	for i, v := range d {
		println(i, v)
	}
	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	newD := map[string]int{}
	err = json.Unmarshal(jsonData, &newD)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range newD {
		println(i, v)
	}
}

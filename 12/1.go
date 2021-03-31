package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p1 := Person{Name: "1"}
	p2 := Person{Name: "2"}
	p3 := Person{"3"}
	persons := []Person{p1, p2, p3}
	fmt.Println(persons)
	for i, _ := range persons {
		persons[i].Name = "123"
	}
	fmt.Println(persons)
}

package main

import "fmt"

func showFullName(s Student) {
	fmt.Println(s.FirstName, s.LastName)
}

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	s1 := Student{
		FirstName: "yerassyl",
		LastName:  "tleugazy",
		Age:       23,
	}
	s2 := Student{
		"asdsda",
		"asdsds",
		32,
	}
	students := []Student{s1, s2}
	for _, v := range students {
		showFullName(v)
	}

}

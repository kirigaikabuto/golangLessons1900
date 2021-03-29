package main

import "fmt"

//struct
type Student struct {
	FirstName string
	LastName  string
	Age       int
	FullName  string
	YearBirth int
}

//methods
func (st *Student) PrintAll() {
	//create call of method CreateFullName()
	//create call of method SetYear()
	fmt.Println(st.FirstName, st.LastName, st.Age, st.FullName)
}

func (st *Student) CreateFullName() {
	st.FullName = st.FirstName + " " + st.LastName
}

//SetYear() method for calculating year of birth (2021-Age)

func main() {
	var st1 Student
	st1 = Student{
		FirstName: "Yerassyl",
		LastName:  "Tleugazy",
		Age:       23,
	}
	st2 := Student{
		FirstName: "123231",
		LastName:  "4123313",
		Age:       123,
	}
	students := []Student{st1, st2}
	for _, v := range students {
		v.PrintAll()
	}
}

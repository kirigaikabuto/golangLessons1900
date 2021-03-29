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
	st.CreateFullName()
	st.SetYear()
	fmt.Println(st.FirstName, st.LastName, st.Age, st.FullName, st.YearBirth)
}

func (st *Student) CreateFullName() {
	st.FullName = st.FirstName + " " + st.LastName
}

func (st *Student) SetYear() {
	st.YearBirth = 2021 - st.Age
}

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

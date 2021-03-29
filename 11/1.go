package main

import "fmt"

//struct
type Student struct {
	FirstName string
	LastName  string
	Age       int
	FullName  string
}

//methods
func (st *Student) PrintAll() {
	fmt.Println(st.FirstName, st.LastName, st.Age, st.FullName)
}

func (st *Student) CreateFullName() {
	st.FullName = st.FirstName + " " + st.LastName
}

func main() {
	var st1 Student
	st1 = Student{
		FirstName: "Yerassyl",
		LastName:  "Tleugazy",
		Age:       23,
	}
	st1.CreateFullName()
	st1.PrintAll()
}

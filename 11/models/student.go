package models

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

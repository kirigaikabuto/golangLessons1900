package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/11/models"

	"os"
)

func main() {
	var st1 models.Student
	st1 = models.Student{
		FirstName: "Yerassyl",
		LastName:  "Tleugazy",
		Age:       23,
	}
	st2 := models.Student{
		FirstName: "123231",
		LastName:  "4123313",
		Age:       123,
	}
	students := []models.Student{st1, st2}
	fmt.Println(students)
	//convert to json
	studentsJson, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	//write data to file
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(studentsJson)
	if err != nil {
		panic(err)
	}
}

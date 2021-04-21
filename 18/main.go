package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

type Student struct {
	Name string `json:"name"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from main 1231231312")
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "87086394516")
}

func CalcPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	a := r.Form.Get("ab")
	b := r.Form.Get("b")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	c := aInt + bInt
	fmt.Println(c)
	st1 := Student{Name: "123231"}
	JsonResponse(w, st1)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	students := []Student{
		{
			Name: "123",
		},
		{
			"456",
		},
		{
			"789",
		},
	}
	JsonResponse(w, students)
}

func JsonResponse(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(value)
}

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/contacts", ContactPage)
	http.HandleFunc("/calculate", CalcPage)
	http.HandleFunc("/students", GetStudents)
	log.Info("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

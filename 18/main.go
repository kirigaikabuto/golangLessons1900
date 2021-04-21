package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

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
		log.Fatal(err)
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	c := aInt + bInt
	fmt.Fprint(w, c)
}

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/contacts", ContactPage)
	http.HandleFunc("/calculate", CalcPage)

	log.Info("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

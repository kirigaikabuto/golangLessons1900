package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateResponse struct {
	Title       string
	Description string
	Price       int
	Error       string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/main.html")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := r.Form.Get("username")
		password := r.Form.Get("password")
		if name == "" {
			http.Error(w, "please write username", http.StatusBadRequest)
			return
		}
		if password == "" {
			http.Error(w, "please write password", http.StatusBadRequest)
			return
		}
		//connect to database and save data
		fmt.Println("form data ", name, password)
		http.ServeFile(w, r, "templates/main.html")
	}
}

func TemplateExample(w http.ResponseWriter, r *http.Request) {
	templateFile, err := template.ParseFiles("templates/template.html")
	if err != nil {
		http.Error(w, "no html file", http.StatusInternalServerError)
		return
	}
	tp := &TemplateResponse{
		Title:       "Golang lesson about template",
		Description: "template example",
		Price:       10000000,
	}
	if r.Method == "GET" {
		tp.Error = ""
		err = templateFile.Execute(w, tp)
		if err != nil {
			http.Error(w, "execute error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		name := r.Form.Get("username")
		password := r.Form.Get("password")
		tp.Error = "Hello mr." + name
		if name == "" {
			tp.Error = "please write username"
		}
		if password == "" {
			tp.Error = "please write password"
		}
		err = templateFile.Execute(w, tp)
		if err != nil {
			http.Error(w, "execute error", http.StatusInternalServerError)
			return
		}
	}

}

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/template", TemplateExample)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

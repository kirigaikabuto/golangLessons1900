package login

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HttpEndpoints interface {
	MainPage(w http.ResponseWriter, r *http.Request)
	TemplateExample(w http.ResponseWriter, r *http.Request)
}

type loginHttpEndpointsStruct struct {
	//connection to database
}

func NewLoginHttpEndpoints(cf Config) HttpEndpoints {
	return &loginHttpEndpointsStruct{}
}

func (lh *loginHttpEndpointsStruct) MainPage(w http.ResponseWriter, r *http.Request) {
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

func (lh *loginHttpEndpointsStruct) TemplateExample(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("data from login ", name, password)
		err = templateFile.Execute(w, tp)
		if err != nil {
			http.Error(w, "execute error", http.StatusInternalServerError)
			return
		}
	}

}

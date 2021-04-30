package movie

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

type HttpEndpoints interface {
	MainPage(w http.ResponseWriter, r *http.Request)
	AddPage(w http.ResponseWriter, r *http.Request)
	AddPageAction(w http.ResponseWriter, r *http.Request)
	DetailPage(idParam string) func(w http.ResponseWriter, r *http.Request)
}

type httpEndpoints struct {
	db MovieStore
}

func NewHttpEndpoints(database MovieStore) HttpEndpoints {
	return &httpEndpoints{db: database}
}

func (h *httpEndpoints) DetailPage(idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars[idParam]
		if !ok {
			http.Error(w, "Please write parameter id", http.StatusInternalServerError)
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templateFile, err := template.ParseFiles("templates/detail_page.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movie, err := h.db.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(movie)
		response := &DetailPageResponse{Movie: *movie}
		err = templateFile.Execute(w, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *httpEndpoints) MainPage(w http.ResponseWriter, r *http.Request) {
	templateFile, err := template.ParseFiles("templates/main_page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	movies, err := h.db.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &MainPageResponse{Movies: movies}
	err = templateFile.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *httpEndpoints) AddPage(w http.ResponseWriter, r *http.Request) {
	response := &AddPageResponse{Error: ""}
	templateFile, err := template.ParseFiles("templates/add_page.html")
	form, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	errorText := form.Get("error")
	fmt.Println("error text", errorText)
	fmt.Println(r.Form.Get("name"))
	if errorText != "" {
		response.Error = errorText
	}
	err = templateFile.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *httpEndpoints) AddPageAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := r.Form.Get("name")
	description := r.Form.Get("description")
	imageUrl := r.Form.Get("image_url")
	errorText := ""
	if len(name) == 0 || len(description) == 0 || len(imageUrl) == 0 {
		errorText = "Please write all data"
		form, _ := url.ParseQuery(r.URL.RawQuery)
		form.Add("error", errorText)
		r.URL.RawQuery = form.Encode()
		http.Redirect(w, r, "/add_page", http.StatusSeeOther)
		return
	}
	movie := &Movie{
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
	}
	movie, err = h.db.Create(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

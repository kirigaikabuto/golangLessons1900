package main

import (
	"fmt"
	"github.com/kirigaikabuto/golangLessons1900/19/login"
	"net/http"
)

func main() {
	loginHttpEndpoints := login.NewLoginHttpEndpoints(login.Config{})
	http.HandleFunc("/", loginHttpEndpoints.MainPage)
	http.HandleFunc("/template", loginHttpEndpoints.TemplateExample)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

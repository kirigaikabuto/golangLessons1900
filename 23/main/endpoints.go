package main

import (
	"encoding/json"
	"github.com/kirigaikabuto/golangLessons1900/23/users"
	"net/http"
)

type HttpEndpoints interface {
	TestEndpoint() func(w http.ResponseWriter, r *http.Request)
}
type httpEndpoints struct {
	//variable connection to db
}

func NewHttpEndpoints() HttpEndpoints {
	return &httpEndpoints{}
}

func (h *httpEndpoints) TestEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := users.User{
			Id:        "1232213",
			Username:  "username123",
			Password:  "asdssadsa",
			FirstName: "1233",
			LastName:  "13213",
			Avatar:    "asdsadsadsadad",
		}
		respondJSON(w, http.StatusOK, user)
		return
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

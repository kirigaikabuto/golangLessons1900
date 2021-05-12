package start

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kirigaikabuto/golangLessons1900/23/users"
	"io/ioutil"
	"net/http"
)

type HttpEndpoints interface {
	TestEndpoint() func(w http.ResponseWriter, r *http.Request)
	TestEndpointWithParam(idParam string) func(w http.ResponseWriter, r *http.Request)
	TestPostEndpoint() func(w http.ResponseWriter, r *http.Request)
	RegisterEndpoint() func(w http.ResponseWriter, r *http.Request)
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

func (h *httpEndpoints) TestEndpointWithParam(idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars[idParam]
		if !ok {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    "Dont have user with that id",
				StatusCode: http.StatusBadRequest,
			})
		}
		var response users.User
		usersData := []users.User{
			{
				Id:        "1",
				Username:  "1",
				Password:  "1",
				FirstName: "1",
				LastName:  "1",
				Avatar:    "1",
			},
			{
				Id:        "2",
				Username:  "2",
				Password:  "2",
				FirstName: "2",
				LastName:  "2",
				Avatar:    "2",
			},
		}
		if idStr == "1" {
			response = usersData[0]
		} else if idStr == "2" {
			response = usersData[1]
		} else {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    "Dont have user with that id",
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		respondJSON(w, http.StatusOK, response)
		return
	}
}

func (h *httpEndpoints) TestPostEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondJSON(w, http.StatusBadRequest, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
			return
		}
		user := &users.User{}
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
		user.Id = "3333"
		respondJSON(w, http.StatusCreated, user)
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

func (h *httpEndpoints) RegisterEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

package main

type HttpEndpoints interface {
}

type httpEndpoints struct {
	//variable connection to db
}

func NewHttpEndpoints() HttpEndpoints {
	return &httpEndpoints{}
}

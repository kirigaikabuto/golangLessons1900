package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func getSum(a, b interface{}) (interface{}, error) {
	var output interface{}
	dataTypeA := reflect.TypeOf(a)
	dataTypeB := reflect.TypeOf(b)
	if dataTypeA.String() == reflect.Float64.String() && dataTypeB.String() == reflect.Float64.String() {
		output = a.(float64) + b.(float64)
	} else if dataTypeA.String() == reflect.Int.String() && dataTypeB.String() == reflect.Int.String() {
		output = a.(int) + b.(int)
	} else if dataTypeA.String() == reflect.String.String() && dataTypeB.String() == reflect.String.String() {
		output = a.(string) + b.(string)
	} else {
		return nil, errors.New("type of data should be the same")
	}
	return output, nil
}

func main() {
	out, err := getSum(3.4, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

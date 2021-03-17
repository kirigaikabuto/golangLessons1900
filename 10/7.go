package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 3
	b := 3.4
	c := "hello"
	d := false
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
}

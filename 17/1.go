package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func main() {
	a := 3
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
	b := "123123"
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)
	var c interface{}
	//c = "1231231212313"
	c = 12323
	//c = &Student{Name: "yerassyl"}
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
}

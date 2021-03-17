package main

import (
	"fmt"
	"reflect"
)

//type название_структуры struct {
//	имя_поля тип_поля(примитивные типы данных или такая же структура)
//	...
//}
//тип данных person у которого есть возраст(age) имя(first_name) фамилия(last_name)

//новый тип данных
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//примитивные типы данных
	//int
	//float64
	//string
	//bool

	//a := 3
	//b := 3.4
	//c := "hello"
	//d := false
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(b))
	//fmt.Println(reflect.TypeOf(c))
	//fmt.Println(reflect.TypeOf(d))
	p := Person{"yerassyl", "tleugazy", 23}
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
	fmt.Println(p.Age)
	fmt.Println(reflect.TypeOf(p))

}

package main

import "fmt"

func printHello() {
	fmt.Println("Hello world")
}

func getSumOfNumbers() {
	a := 3
	b := 4
	c := a + b
	fmt.Println(c)
}

func HelloYerassyl() {
	fmt.Println("Hello Yerassyl")
}

func main() {
	//функция и процедура
	//процедура это неких блок кода который выполняется только тогда когда к нему обратятся по имени процедуры
	printHello()
	getSumOfNumbers()
	HelloYerassyl()
}

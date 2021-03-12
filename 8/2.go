package main

import "fmt"

//аргументы -> это входные параметры которые необходимы для работы процедуры или фукнции
func printData(word string) {
	fmt.Println(word)
}

func getSum(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	printData("Hello yerassyl")
	getSum(3, 4)
}

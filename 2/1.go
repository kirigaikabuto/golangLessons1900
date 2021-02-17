package main

import "fmt"

func main() {
	var a, b int
	//%d - decimal -> десятичные цифры
	//%s - string -> строки
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	c := a + b
	fmt.Println(c)
}

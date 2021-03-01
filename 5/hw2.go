package main

import "fmt"

//2) ряд фибоначчи
//вводится какое либо число и небходимо посчитать число фибоначи до этого числа
//n=10
//0 1 1 2 3 5 8 13
func main() {
	var n int
	fmt.Scanf("%d", &n)
	start1 := 0
	start2 := 1
	for true {
		c := start1 + start2
		if c > n {
			break
		}
		fmt.Println(c)
		tmp := start1
		start1 = start2
		start2 = start1 + tmp
	}
}

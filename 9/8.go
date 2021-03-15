package main

import "fmt"

func main() {
	a := [][]int{}
	for i := 0; i < 4; i++ {
		arr := []int{1, 1, 1, 1}
		a = append(a, arr)
	}
	for _, v := range a {
		fmt.Println(v)
	}
}

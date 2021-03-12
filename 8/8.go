package main

import "fmt"

func replaceElements(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			a[i] = 0
		}
	}
	return a
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := replaceElements(arr)
	fmt.Println(arr)
	fmt.Println(arr2)
}

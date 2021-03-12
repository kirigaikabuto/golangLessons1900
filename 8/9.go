package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	arr1 := arr
	arr[0] = 123
	fmt.Println(arr)
	fmt.Println(arr1)
}

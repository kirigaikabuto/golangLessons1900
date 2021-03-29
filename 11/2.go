package main

import "fmt"

func addElement(arr []int) {
	arr[0] = 123
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	addElement(arr)
	fmt.Println(arr)
}

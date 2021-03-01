package main

import "fmt"

func main() {
	arr := []int{1, 2, 33, 4, 5}
	//values 1 2 33 4 5
	//indexs 0 1 2  3 4

	//update
	arr[0]=123
	//get
	print(arr[0])
	fmt.Println(arr[4])
}

package main

import "fmt"

func checkForEven(k int) bool {
	if k%2 == 0 {
		return true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		if checkForEven(arr[i]) {
			fmt.Println(arr[i])
		}
	}
}

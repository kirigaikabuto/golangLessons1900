package main

import "fmt"

func main() {
	arr := [5]int{-3, -3, -4, -1, -5}
	n := len(arr)
	maxi := arr[0]
	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	fmt.Println(maxi)
}

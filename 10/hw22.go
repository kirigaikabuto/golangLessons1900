package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	arr = append(arr, arr[0])
	arr = append(arr, arr[1])
	fmt.Println(arr)
	n := len(arr)
	sumi := 0
	maxi := 0
	for i := 0; i < n-2; i++ {
		sumi = arr[i] + arr[i+1] + arr[i+2]
		if sumi > maxi {
			maxi = sumi
		}
	}
	fmt.Println(maxi)
}

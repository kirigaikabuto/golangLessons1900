package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	n := len(arr)
	sumi := 0
	maxi := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			sumi = arr[i] + arr[0] + arr[1]
		} else if i == n-2 {
			sumi = arr[i] + arr[i+1] + arr[0]
		} else {
			sumi = arr[i] + arr[i+1] + arr[i+2]
		}
		if sumi > maxi {
			maxi = sumi
		}
	}
	fmt.Println(maxi)
}
// 1
//4 2
// 3

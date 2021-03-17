package main

import "fmt"

func main() {
	arr2 := [][]int{
		{25, 1},
		{70, 1},
		{100, 0},
		{3, 1},
	}
	max := arr2[0][0]
	index := 0
	for i := 0; i < len(arr2); i++ {
		if arr2[i][0] > max && arr2[i][1] == 1 {
			max = arr2[i][0]
			index = i + 1
		}
	}
	fmt.Println(index)
}

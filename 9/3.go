package main

import "fmt"

func main() {
	//одномерный slice
	//arr := []int{1, 2, 3, 4, 5}
	//fmt.Println(arr[0])
	//   0 1 2
	//0  1 2 3
	//1  4 5 6
	//2  7 8 9
	//3 на 3
	twoArr := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	fmt.Println(len(twoArr))
	//1
	//4
	//8
	//9
}

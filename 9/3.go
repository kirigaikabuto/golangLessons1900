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
	twoArr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(twoArr[1][1])
	fmt.Println(twoArr[0][2])
}

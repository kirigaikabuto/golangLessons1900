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
		{1, 2, 3, 5}, {4, 5, 6, 0}, {7, 8, 9, 1}, {1, 2, 3, 5},
	}
	sumi := 0
	for i := 0; i < len(twoArr); i++ {
		for j := 0; j < len(twoArr[i]); j++ {
			sumi += twoArr[i][j]
		}
	}
	fmt.Println(sumi)
	//1
	//4
	//8
	//9
}

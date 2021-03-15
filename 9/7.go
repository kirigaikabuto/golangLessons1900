package main

import "fmt"

func main() {
	twoArr := [][]int{
		{1, 2, 3, 5},
		{4, 5, 6, 0},
		{7, 8, 9, 1},
		{1, 2, 3, 5},
	}
	//for i := 0; i < len(twoArr); i++ {
	//	fmt.Println("1")
	//	for j := 0; j < len(twoArr[i]); j++ {
	//		fmt.Println("2")
	//	}
	//}
	sumi := 0
	for i, v := range twoArr {
		for l, k := range v {
			sumi += k
			twoArr[i][l] = 0
		}
	}
	fmt.Println(sumi)
	fmt.Println(twoArr)
}

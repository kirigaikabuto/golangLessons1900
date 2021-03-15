package main

import "fmt"

func main() {
	//arr := [][]int{
	//	{25, 1},
	//	{70, 1},
	//	{100, 0},
	//	{3, 1},
	//}
	arr := [][]int{}
	var n int
	fmt.Scanf("%d \n", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d \n", &a, &b)
		nestedArr := []int{a, b}
		arr = append(arr, nestedArr)
	}
	fmt.Println(arr)
}

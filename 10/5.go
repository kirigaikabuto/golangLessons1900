package main

import (
	"fmt"
	"math/rand"
)

func CreateTwoDimensionalArray(rows, columns, randomRange int) [][]int {
	arr := [][]int{}
	//code
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			k := rand.Intn(randomRange)
			a = append(a, k)
		}
		arr = append(arr, a)
	}
	return arr
}

func main() {
	arr := CreateTwoDimensionalArray(10, 20, 15)
	for _, v := range arr {
		fmt.Println(v)
	}
	for i := 0; i < len(arr); i++ {
		shetchik := 0
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 5 {
				shetchik += 1
			}
		}
		if shetchik >= 3 {
			fmt.Println("Повторение в массиве = ", shetchik, "", "Строка = ", i)
		}
	}
}

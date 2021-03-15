package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createTwoDimensionalArray(rows, columns, randomRange int) [][]int {
	arr := [][]int{}
	//code
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			//create random number
			k := rand.Intn(randomRange)
			//add elements to the nested array
			a = append(a, k)
		}
		//add array into arr array
		arr = append(arr, a)
	}
	return arr
}

func CreateArray(n, randomRange int) []int {
	//тут мы генерируем числа для рандома что он заполнился числами
	rand.Seed(time.Now().UnixNano())
	//выдаст любое число до 120
	arr := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(randomRange)
		arr = append(arr, k)
	}
	return arr
}

func main() {
	arr := createTwoDimensionalArray(3, 4, 20)
	for _, v := range arr {
		fmt.Println(v)
	}
}

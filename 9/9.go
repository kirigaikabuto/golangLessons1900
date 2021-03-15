package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	arr := CreateArray(20, 10)
	fmt.Println(arr)
}

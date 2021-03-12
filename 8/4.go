package main

import (
	"fmt"
	"math"
)

//func названии_функции(параметры...) тип_элемента_который_будет_возращен  {
//блок кода
//return то_что_мы_хотим_вернуть_обратно_откуда_вызывали
//}
func getMax(a []int) int {
	maxi := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
		}
	}
	return maxi
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 2, 323, 23, 4, 232}
	arr3 := []int{1, 23, 4, 123, 23213, 43434, 4545, 12321, 1313213, 343443, 54565}
	maxi1 := getMax(arr1)
	maxi2 := getMax(arr2)
	maxi3 := getMax(arr3)
	fmt.Println(maxi1, maxi2, maxi3)
	k := math.Pow(float64(3),float64(2))
	fmt.Println(k)
}

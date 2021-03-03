package main

import "fmt"

//найти в массиве четные и нечетные элементы и
//поместить их в разные массивы и вывести эти массивы
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	n := len(arr1)
	sumi := 0
	var arr2 [100]int
	var arr3 [100]int
	for i := 0; i < n; i++ {
		sumi = sumi + arr1[i]
	}
	avg := sumi / n
	i2 := 0
	i3 := 0
	for i := 0; i < n; i++ {
		if arr1[i] < avg {
			arr2[i2] = arr1[i]
			i2++
		}
		if arr1[i] >= avg {
			arr3[i3] = arr1[i]
			i3++
		}
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

package main

import "fmt"

//процедура или функция убирает повторяющий кусок кода и делает его общим для всех
func printArr(a []int, name string) {
	fmt.Println(name)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
//array1
//4 1
//array2
//323 2
func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 2, 323, 23, 4, 232}
	arr3 := []int{1, 23, 4, 123, 23213, 43434, 4545, 12321, 1313213, 343443, 54565}
	printArr(arr1, "array1")
	printArr(arr2, "array2")
	printArr(arr3, "array3")
}

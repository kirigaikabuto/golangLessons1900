package main

import "fmt"

func getMaxAndMin(a []int) (int, int) {
	maxi := a[0]
	mini := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
		}
		if a[i] < mini {
			mini = a[i]
		}
	}
	return maxi, mini
}

func getAvg(a []int) int {
	sumi := 0
	avg := 0
	for i := 0; i < len(a); i++ {
		sumi += a[i]
	}
	avg = sumi / len(a)
	return avg
}

func printInfo(a []int) {
	maxi, mini := getMaxAndMin(a)
	avg := getAvg(a)
	fmt.Println("maxi =", maxi)
	fmt.Println("mini =", mini)
	fmt.Println("avg =", avg)
	fmt.Println("")
}

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{4, 5, 6, 2, 323, 23, 4, 232}
	arr3 := []int{1, 23, 4, 123, 23213, 43434, 4545, 12321, 1313213, 343443, 54565}
	printInfo(arr1)
	printInfo(arr2)
	printInfo(arr3)
}

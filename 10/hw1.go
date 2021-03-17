package main

import "fmt"

//https://acmp.ru/index.asp?main=task&id_task=264
//готов
func main() {
	arr := [10]int{10, 20, 30, 1, -10, 1, 2, 3}
	max := 0
	ottepel := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] >= 0 {
			ottepel += 1
		} else {
			if ottepel > max {
				max = ottepel
			}
			ottepel = 0
		}
	}
	fmt.Println("Продолжительная оттепель = ", max)
}

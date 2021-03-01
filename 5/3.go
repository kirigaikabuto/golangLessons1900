package main

import "fmt"

func main() {
	//вводит сперва n-> количество чисел в массиве
	//затем поочереди каждое число вводится
	var n int
	fmt.Scanf("%d", &n)
	arr := [100]int{}
	fmt.Println(arr)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	fmt.Println(arr)
	k := len(arr)
	fmt.Println(k)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi = sumi + arr[i]
	}
	fmt.Println(sumi)
}

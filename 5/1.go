package main

import "fmt"

func main() {
	//a := 3
	//b := "hell world"
	//c := 3.15
	//d := false
	//массив
	var arr [5]int
	//values 0 0 0 0 0
	//index  0 1 2 3 4
	arr[0] = 1
	arr[1] = 2
	arr[2] = 13
	arr[3] = 10
	arr[4] = 3
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	sumi := arr[0] + arr[1]
	fmt.Println(sumi)

}

package main

import "fmt"

func main() {
	//диапозон чисел вводится с клавиатуры и необходмо узнать количество
	//четных и неченых в этом промежутке
	//1 7
	//5 четных
	//5 нечетных
	znak := ""
	for true {
		fmt.Println("Hello friend")
		fmt.Scanf("%s", &znak)
		if znak == "0" {
			break
		}
	}
}

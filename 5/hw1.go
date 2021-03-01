package main

import "fmt"

//1) калькулятор
//+ / - *
//
//программа запрашивает вид операции (+,-,/,*,0)
//если вид операции это 0 значит нужнов выйти из программы
//но а если знак операции (+,-,/,*)
//вид операции:+
//первое число:3
//второе число:4
//Ответ:7
//вид операции:
func main() {
	for true {
		var operation string
		fmt.Scanf("%s", &operation)
		if operation == "0" {
			break
		} else if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
			var a, b int
			fmt.Scanf("%d", &a)
			fmt.Scanf("%d", &b)
			if operation == "+" {
				fmt.Println(a + b)
			}
			if operation == "*" {
				fmt.Println(a * b)
			}
		} else {
			fmt.Println("there is no this operation")
		}
	}
}

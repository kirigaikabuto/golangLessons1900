package main

import "fmt"

func main() {
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
	a := ""
	var b, c int
	for true {
		fmt.Scanf("%s", &a)
		fmt.Printf("вид операции: %s \n", a)
		if a == "0" {
			break
		}
		if a == "+" || a == "-" || a == "/" || a == "*" {
			fmt.Scanf("%d %d", &b, &c)
			fmt.Printf("первое число: %d \n", b)
			fmt.Printf("второе число: %d \n", c)
			if a == "+" {
				fmt.Printf("Ответ: %d", b+c)
			}
			if a == "-" {
				fmt.Printf("Ответ: %d \n", b-c)
			}
			if a == "/" {
				if c != 0 {
					fmt.Printf("Ответ: %d \n", b/c)
				} else {
					fmt.Println("Деление на ноль!")
				}
			}
			if a == "*" {
				fmt.Printf("Ответ: %d \n", b*c)
			}
		}

	}
}

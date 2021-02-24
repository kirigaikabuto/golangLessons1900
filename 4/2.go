package main

import "fmt"

func main() {
	//i=0 счетчик
	//i<5
	//i=i+1 -> i+=1 -> i++
	//for созадем счетчик; условия работы счетчика; условия изменения счетчика{
	//	действие1
	//	действие2
	//}
	//i=0 +
	//i=2 +
	//i=4 +
	//i=6
	fmt.Println("for start")
	for i := 0; i < 5; i+=2 {
		fmt.Println("Yerassyl")
	}
	fmt.Println("for end")

}

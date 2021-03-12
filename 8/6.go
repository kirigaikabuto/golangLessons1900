package main

import "fmt"

func main() {
	//int - одно число
	//[]int - массив из чисел
	//char - один символ
	//string -> массив из символол
	var email string
	fmt.Scanf("%s", &email)
	isCorrect := false
	for i := 0; i < len(email); i++ {
		//email[i] -> t,l,e,
		if email[i] == '@' {
			isCorrect = true
			break
		}
	}
	if isCorrect {
		fmt.Println("ok")
	} else {
		fmt.Println("should contain @")
	}
}

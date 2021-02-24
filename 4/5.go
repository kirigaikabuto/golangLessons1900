package main

import "fmt"

func main() {
	//необходимо посчитать сумму цифр 0-20 используя цикл
	//0+1+2+3+4+5.... =
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	//5! -> 1*2*3*4*5 = 120
	//5
	//120

}

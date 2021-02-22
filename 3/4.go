package main

import "fmt"

func main() {
	var salary, month, price int
	fmt.Scanf("%d", &salary)
	fmt.Scanf("%d", &month)
	fmt.Scanf("%d", &price)
	r := float64(price / month)
	s := float64(salary) * 0.5
	if s > r {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

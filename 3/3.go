package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	c := a + b
	if c == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}
}

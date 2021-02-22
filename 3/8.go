package main

import "fmt"

func main() {
	a := 10
	b := 22
	c := 5
	if a > b && a > c {
		fmt.Println(a)
	}
	if b > a && b > c {
		fmt.Println(b)
	}
	if c > a && c > b {
		fmt.Println(c)
	}
}

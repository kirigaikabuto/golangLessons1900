package main

import "fmt"

func main() {
	//(a+b)*(b+c)
	//a := 3
	//b := 4
	//c := 5
	a, b, c := 3, 4, 5
	d := a + b
	e := b + c
	f := d * e
	fmt.Println(f)
}

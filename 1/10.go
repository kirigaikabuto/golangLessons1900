package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 10
	c := 10
	perimeter := a + b + c
	p := perimeter / 2
	d := p * (p - a) * (p - b) * (p - c)
	s := math.Sqrt(float64(d))
	fmt.Println(perimeter)
	fmt.Println(s)
}

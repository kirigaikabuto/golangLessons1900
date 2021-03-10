package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 45}
	b := []int{}
	c := []int{}
	a = append(a, 3)
	fmt.Println(a[3])
	a = append(a, 123)
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			b = append(b, a[i])
		} else {
			c = append(c, a[i])
		}
	}
	fmt.Println(b)
	fmt.Println(c)
}

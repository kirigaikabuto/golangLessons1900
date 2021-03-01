package main

import "fmt"

func main() {
	var n, a int
	sumi := 0
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		sumi = sumi + a
	}
	fmt.Println(sumi)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := [100]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	maxi := 0
	//20 30 40 50 60
	//20 30 -> 30
	//30 40 -> 40
	//40 50 -> 50
	//50 60 -> 60
	//20 0 -> maxi = 20
	//30 20 -> maxi = 30
	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	fmt.Println(maxi)
}

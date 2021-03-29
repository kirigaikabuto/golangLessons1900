package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		panic(err)
	}
	name := "Yerassyl"
	n, err := file.WriteString(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

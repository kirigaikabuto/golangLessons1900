package main

import "fmt"

func PrintData(s string) {
	fmt.Println(s)
}

func main() {
	PrintData("sdsdss")
	SendMessage := func(l string) {
		fmt.Println(l)
	}
	SendMessage("hello world")
}

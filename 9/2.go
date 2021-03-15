package main

import "fmt"
import "github.com/kirigaikabuto/golangLessons1900/9/utils"



func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newArr := utils.GetOnlyEven(arr)
	fmt.Println(newArr)
}

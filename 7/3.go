package main

import "fmt"

func main() {
	//вводится количество элементов в массиве
	//вводится массив из элементов
	//вводится число которое необходимо найти и если оно есть в массиве вывести есть иначе нет
	//arr := []int{1, 11, 2, 1, 3, 4, 2, 2}
	//1 - 2
	//11 -1
	//2 - 3
	//3 - 1
	//1 - 1
	//4 - 1
	arr := []int{1, 11, 2, 1, 3, 4, 2, 2}
	numbers := []int{}
	repeats := []int{}
	//цикл для того чтобы взять каждое число по очереди
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		counter := 0
		//цикл для того чтобы число которое у нас идет по очереди сравнить со всеми числами в массиве
		for j := 0; j < len(arr); j++ {
			if temp == arr[j] {
				counter += 1
			}
		}
		isNotExist := true
		//цикл для того чтобы понять не существует ли уже цифры которой мы идем по очереди в массиве где мы храним числа которые уже отработали
		for k := 0; k < len(numbers); k++ {
			if numbers[k] == temp {
				isNotExist = false
				break
			}
		}
		//если мы не нашли число которое у нас идет по очерди то тогда мы добавляем его и количество его повторений
		if isNotExist {
			numbers = append(numbers, temp)
			repeats = append(repeats, counter)
		}
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i], repeats[i])
	}
}

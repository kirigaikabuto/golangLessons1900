package main

func main() {
	a := 16
	if a%8 == 0 {
		println("остатка при деление на 8 нету")
	} else if a%6 == 0 {
		println("остатка при деление на 6 нету")
	} else if a%4 == 0 {
		println("остатка при деление на 4 нету")
	} else if a%2 == 0 {
		println("остатка при деление на 2 нету")
	} else {
		println("остаток есть всегда")
	}

}

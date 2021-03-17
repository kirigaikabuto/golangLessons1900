package main

import "fmt"

type Book struct {
	Name  string
	Price int
}

func(b Book) printBook() {
	fmt.Println("Название книги:", b.Name, "Цена:", b.Price)
}

func main() {
	b1 := Book{
		Name:  "book1",
		Price: 2,
	}

	b2 := Book{
		Name:  "book2",
		Price: 3,
	}

	b3 := Book{
		Name:  "book2",
		Price: 1,
	}

	books := []Book{b1, b2, b3}
	maxi := books[0]
	for _, v := range books {
		if v.Price > maxi.Price {
			maxi = v
		}
	}
	b1.printBook()
	b2.printBook()
	b3.printBook()
	maxi.printBook()
}

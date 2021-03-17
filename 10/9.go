package main

import "fmt"

type Book struct {
	Name  string
	Price int
}

func (b Book) printBook() {
	fmt.Println("Название книги:", b.Name, "Цена:", b.Price)
}

type Library struct {
	Name    string
	Address string
	Books   []Book
}

func (l Library) getExpensiveBook() {
	books := l.Books
	maxi := books[0]
	for _, v := range books {
		if v.Price > maxi.Price {
			maxi = v
		}
	}
	maxi.printBook()
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
	l1 := Library{
		Name:    "Library1",
		Address: "Library1",
		Books:   books,
	}
	l1.getExpensiveBook()
}

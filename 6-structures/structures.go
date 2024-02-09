package main

import "fmt"

type Book struct {
	name   string
	page   int
	isRead bool
}

func main() {
	var book1 = Book{"metamorphosis", 98, false}
	fmt.Println(book1)
	book1.isRead = true
	fmt.Println(book1)
}

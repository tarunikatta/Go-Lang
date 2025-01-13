package main

import (
	"book/book"
	"fmt"
)

func main() {
	b := book.Book{Title: "The Alchemist", Author: "Paulo Coelho", Pages: 150, CopiesAvailable: 10}
	k := book.Book{Title: "The Hobbit", Author: "J.R.R. Tolkien", Pages: 250, CopiesAvailable: 5}

	b.Display()
	fmt.Println(b.Borrow())
	b.Display()
	b.ReturnBook()
	b.Display()

	book.SwapTitles(&b, &k)
	fmt.Println("\nAfter swapping titles:")
	b.Display()
	k.Display()
}

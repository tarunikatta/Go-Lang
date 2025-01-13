package book

import "fmt"

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func (b Book) Display() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nCopies Available: %d\n", b.Title, b.Author, b.Pages, b.CopiesAvailable)
}

func (b *Book) Borrow() string {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable--
		return "Borrow successful!"
	}
	return "No copies available to borrow."
}

func (b *Book) ReturnBook() {
	b.CopiesAvailable++
}

func SwapTitles(b1, b2 *Book) {
	b1.Title, b2.Title = b2.Title, b1.Title
}

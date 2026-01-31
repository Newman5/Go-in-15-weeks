package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	b := Book{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		Pages:  464,
	}
	fmt.Println(describeBook(b))
}

func describeBook(b Book) string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

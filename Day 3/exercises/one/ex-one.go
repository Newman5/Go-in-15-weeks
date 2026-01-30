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
	fmt.Println(b)
	c := addPagesValue(b, 10)
	fmt.Println(&c)
	fmt.Println(&b)
	fmt.Println(addPagesPointer(&b, 10))

}

func addPagesValue(b Book, n int) Book {
	b.Pages += n
	return b
}

func addPagesPointer(b *Book, n int) *Book {
	b.Pages += n
	return b
}

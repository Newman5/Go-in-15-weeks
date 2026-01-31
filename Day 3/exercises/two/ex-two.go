package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	books := []Book{
		{Title: "A", Pages: 100},
		{Title: "B", Pages: 200},
	}

	fmt.Println(updatePages(books, "A", 150))
	fmt.Println(books)

}

// func updatePages(books *[]Book, title string, pages int) (*[]Book, error) {
// 	if len(*books) == 0 {
// 		return nil, fmt.Errorf("books is empty")
// 	}

// 	for i := range *books {
// 		if (*books)[i].Title == title {
// 			(*books)[i].Pages = pages
// 			return books, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("book not found")
// }

func updatePages(books []Book, title string, pages int) error {
	if len(books) == 0 {
		return fmt.Errorf("books is empty")
	}
	for i := range books {
		if books[i].Title == title {
			books[i].Pages = pages
			return nil
		}
	}
	return fmt.Errorf("book not found")
}

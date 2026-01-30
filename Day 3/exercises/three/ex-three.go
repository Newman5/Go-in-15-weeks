package main

import "fmt"

type Item struct {
	ID    int
	Name  string
	Count int
}

func main() {
	items := []Item{
		{ID: 1, Name: "Item 1", Count: 10},
		{ID: 2, Name: "Item 2", Count: 20},
		{ID: 3, Name: "Item 3", Count: 30},
	}
	result1 := updateCount(items, 1, 10)
	fmt.Println(items, result1)

	result2 := updateCount(items, 4, 10)
	fmt.Println(items, result2)
}

func updateCount(items []Item, id int, delta int) error {
	if len(items) == 0 {
		return fmt.Errorf("items is empty")
	}
	for i := range items {
		if items[i].ID == id {
			items[i].Count += delta
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

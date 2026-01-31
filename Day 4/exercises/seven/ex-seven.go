// Task
// Build a small program that:
// Stores item names and quantities using a map
// Supports:
// add item
// update quantity
// remove item
// print all items
// Uses:
// functions
// maps
// len
// delete
// range
// comma-ok where appropriate
package main

import "fmt"

type Item struct {
	Name     string
	Quantity int
}

func main() {
	items := make(map[string]Item)
	items["a"] = Item{Name: "a", Quantity: 10}
	items["b"] = Item{Name: "b", Quantity: 20}
	items["c"] = Item{Name: "c", Quantity: 30}
	fmt.Println(items)
	addItem(items, "d", 40)
	fmt.Println(items)
	updateQuantity(items, "d", 50)
	fmt.Println(items)
	removeItem(items, "d")
	fmt.Println(items)
	printAllItems(items)
}

func addItem(items map[string]Item, name string, quantity int) error {
	_, ok := items[name]
	if !ok {
		items[name] = Item{Name: name, Quantity: quantity}
		return nil
	}
	return fmt.Errorf("item already exists")

}

func updateQuantity(items map[string]Item, name string, quantity int) error {
	_, ok := items[name]
	if ok {
		removeItem(items, name)
		addItem(items, name, quantity)
		return nil
	}
	return fmt.Errorf("item not found")
}

func removeItem(items map[string]Item, name string) {
	delete(items, name)
}

func printAllItems(items map[string]Item) {
	for k, v := range items {
		fmt.Println(k, v)
	}
}

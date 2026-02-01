// Exercise 3 — Method Sets (edge case)
// Goal: Understand addressability.
// Keep Counter from Exercise 2.
// Call IncrementPtr() on:
// Counter{Value: 10}
// Observe the compiler error.
// Fix it without changing the method.

package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) IncrementPtr() {
	c.Value++
}

func main() {
	counter := Counter{Value: 10}
	// Counter{Value: 10}.IncrementPtr()
	counter.IncrementPtr()
	fmt.Println(counter.Value)
}

// Exercise 2 — Value vs Pointer Receiver (mutation test)
// Goal: Lock in mutation rules.
// Define:
// type Counter struct {
// 	Value int
// }
// Implement:
// Increment() with a value receiver
// IncrementPtr() with a pointer receiver
// In main:
// Call both
// Print the value after each call

package main

import "fmt"

type Counter struct {
	Value int
}

func Increment(c Counter) {
	c.Value++
}

func IncrementPtr(c *Counter) {
	c.Value++
}

func main() {
	c := Counter{Value: 0}
	Increment(c)
	fmt.Println(c.Value)
	IncrementPtr(&c)
	fmt.Println(c.Value)
}

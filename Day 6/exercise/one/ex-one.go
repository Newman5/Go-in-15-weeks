// Exercise 1—Goroutine lifetime
// Goal: Prove goroutines don’t keep programs alive.
// Write a program that:
// Starts a goroutine that prints after 1 second
// main exits immediately
// Observe output
// Fix it using a channel

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello, World!")
		// Notice how "Hello, World!" is not printed
	}()
}

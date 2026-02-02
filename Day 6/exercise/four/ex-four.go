// Exercise 4 — Deadlock
// Goal: Trigger and understand deadlock.
// Use an unbuffered channel
// Send and receive in main (no goroutine)
// Run and observe runtime error
// Fix by adding go

package main

import "fmt"

func main() {
	// ch := make(chan int) // This will cause a deadlock and an error because the main go routine is blocked after the channel receives the value
	// ch := make(chan int, 1) // This will pass the deadlock
	// Similarly this will also pass the deadlock by creating a sub go routine
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

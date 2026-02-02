// Exercise 2 — Unbuffered synchronization
// Goal: Observe handshake behavior.
// Create an unbuffered channel
// Start a goroutine that:
// Sends a value
// Prints "after send"
// In main, receive and print value
// Confirm "after send" prints only after receive

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		time.Sleep(1 * time.Second)
		fmt.Println("after send")
	}()

	fmt.Println("before receive")
	fmt.Println(<-ch)
	time.Sleep(2 * time.Second) // Wait for the goroutine to finish

}

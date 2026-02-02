// Exercise 5 — Close + range
// Goal: Producer ownership.
// Goroutine sends 3 values then closes channel
// main ranges over channel and prints values
// Remove close and observe behavior
// Restore close

package main

import "fmt"

func main() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// 	close(ch) // If channel is not closed then range will block and produce error
	// }()

	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // If channel is not closed then range will block and produce error
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

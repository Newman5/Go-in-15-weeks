// Exercise 7 — Directional channels
// Goal: Compiler enforcement.
// Write:
// func send(ch chan<- int)
// func recv(ch <-chan int)
// Try to violate direction rules
// Observe compiler error

package main

import "fmt"

func send(ch chan<- int) {
	ch <- 1
	// fmt.Println(<-ch)
}

func recv(ch <-chan int) {
	fmt.Println(<-ch)
	// ch <- 1
}

func main() {
	ch := make(chan int)
	go send(ch)
	// go recv(ch)
	recv(ch)

}

// Exercise 6 — comma ok
// Goal: Detect closure.
// Receive twice from a channel closed after one send
// Print (value, ok) both times
// Confirm zero value + ok == false

package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
}

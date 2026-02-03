// Exercise (do this):
// Create two channels: data and done.
// Start one goroutine that sends data after a short delay.
// Start another goroutine that closes done earlier.
// Use select to react to whichever happens first.
// Run multiple times and observe behavior.

package main

import (
	"fmt"
)

func main() {
	data := make(chan int)
	done := make(chan bool)

	// var wg sync.WaitGroup
	// wg.Add(2)
	go func() {
		// defer wg.Done()
		// time.Sleep(1 * time.Second)
		data <- 1
	}()

	go func() {
		// defer wg.Done()
		done <- true
		// close(done)
	}()

	// When case <-done, the channel already set the value, so the printline in the case will show empty channel
	// for range 2 {
	// 	select {
	// 	case <-data:
	// 		fmt.Println(<-data)
	// 	case <-done:
	// 		fmt.Println(<-done)
	// 	}
	// }

	// wg.Wait()

	for range 2 {
		select {
		case v1 := <-data:
			fmt.Println(v1)
		case v2 := <-done:
			fmt.Println(v2)
		}
	}
	fmt.Println("All done")
}

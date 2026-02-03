// Exercise (do this):
// Write a program that launches 3 goroutines.
// Each goroutine prints its ID after doing some work (e.g., a loop).
// Ensure main waits correctly using sync.WaitGroup.
// Remove wg.Wait() and observe what changes.
// Add it back and confirm deterministic completion.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("ID: 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("ID: 2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("ID: 3")
	}()

	wg.Wait()
	fmt.Println("All Routines are Done")
}

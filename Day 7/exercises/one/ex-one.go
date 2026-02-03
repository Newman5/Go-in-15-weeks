// Exercise
// Write a program where 3 goroutines increment the same counter 1000 times each.
// Run it without a mutex.
// Observe the incorrect final value.
// Add a mutex.
// Run again and confirm correctness.
// Run both versions with -race.

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3000)

	for range 3000 {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait() // Without waitGroups, the last print line is executed immediately after the for loop. No guarentee that all go routines are finished.

	fmt.Println(counter)
}

package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int)
// 	count := sync.WaitGroup{}
// 	count.Add(2)
// 	go func() {
// 		defer count.Done()
// 		ch <- 1 * 3 * 7 * 9 * 11
// 		ch <- 2
// 	}()
// 	go func() {
// 		defer count.Done()
// 		fmt.Println(<-ch)
// 		fmt.Println(<-ch)
// 	}()

// 	count.Wait()
// 	fmt.Println("done")
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	count := sync.WaitGroup{}
// 	count.Add(2)
// 	go func() {
// 		defer count.Done()
// 		ch1 <- 1
// 		fmt.Println(<-ch2)
// 	}()
// 	go func() {
// 		defer count.Done()
// 		fmt.Println(<-ch1)
// 		ch2 <- 2
// 		// fmt.Println(<-ch1)
// 	}()
// 	count.Wait()
// 	fmt.Println("Will it block?")

// }

// func main() {
// 	done := make(chan bool) // 1. Create a signaling channel

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(time.Second)

// 		done <- true // 2. Send signal when finished
// 		close(done)
// 	}()

// 	result, ok := <-done // 3. Block here until a value is received
// 	fmt.Println(result, ok)
// 	result, ok = <-done
// 	fmt.Println(result, ok)
// 	fmt.Println("All done! Main exiting.")
// }

// func main() {
// 	ch := make(chan int, 3)
// 	count := sync.WaitGroup{}
// 	count.Add(2)
// 	go func() {
// 		defer count.Done()
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 		close(ch)
// 	}()

// 	go func() {
// 		defer count.Done()
// 		for range 4 {
// 			v, ok := <-ch
// 			fmt.Println(v, ok)
// 		}
// 	}()

// 	count.Wait()
// 	fmt.Println("All done")
// }

// type Counter struct {
// 	// mu    sync.Mutex
// 	Value int
// }

// func (c *Counter) Increment() {
// 	// c.mu.Lock()
// 	c.Value++
// 	// c.mu.Unlock()
// }

// func main() {
// 	counter := Counter{}
// 	count := sync.WaitGroup{}
// 	count.Add(1000)
// 	for range 1000 {
// 		go func() {
// 			defer count.Done()
// 			counter.Increment()
// 		}()
// 	}
// 	count.Wait()
// 	fmt.Println(counter.Value)
// }

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()
	go func() {
		ch <- 3
	}()
	for range 1 {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(1 * time.Second):
			fmt.Println("Default")
			return
		}
	}

}

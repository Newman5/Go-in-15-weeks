// Exercise 8 — Integration (Day 3 + Day 5)
// Goal: Combine pointers, structs, interfaces.
// Create a struct with internal state:
// type Store struct {
// 	Count int
// }
// Implement:
// Increment() (pointer receiver)
// Notify(string) (value or pointer — you choose)
// Make Store satisfy Notifier.
// Pass it into Send.

package main

import "fmt"

type Incrementer interface {
	Increment()
}

type Store struct {
	Count int
}

func (s *Store) Increment() {
	s.Count++
}

// func (s Store) Increment() {
// 	s.Count++
// }

func main() {
	store := Store{Count: 0}
	store.Increment()
	fmt.Println(store.Count)
	Send(&store)
	// Below Send function will give error unless we change the Increment method to a value receiver
	// Send(store)
	fmt.Println(store.Count)
}

func Send(n Incrementer) {
	n.Increment()
}

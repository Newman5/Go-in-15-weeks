// Exercise 6 — Interface Boundary (why interfaces exist)
// Goal: Make the testing use-case real.
// Write:
// func Send(n Notifier) {
// 	n.Notify("test")
// }
// Create:
// type FakeNotifier struct {
// 	Called bool
// }
// Implement Notify to flip Called = true.
// Call Send(&FakeNotifier{}).
// Verify Called == true.

package main

import "fmt"

type Notifier interface {
	Notify(msg string)
}

type FakeNotifier struct {
	Called bool
}

func (f *FakeNotifier) Notify(msg string) {
	f.Called = true
}

func main() {
	notifier := FakeNotifier{Called: false}
	notifier.Notify("test")
	fmt.Println(notifier.Called)
}

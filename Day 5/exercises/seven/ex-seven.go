// Exercise 7 — The nil Interface Trap (advanced but necessary)
// Goal: See the bug before it bites you.
// Write:
// var f *FakeNotifier = nil
// var n Notifier = f
// Print:
// fmt.Println(n == nil)
// Attempt:
// n.Notify("hi")
// Observe and explain:
// Why the comparison is false
// Why the call panics

package main

import "fmt"

type Notifier interface {
	Notify(msg string)
}

type FakeNotifier struct {
	Called bool
}

func (f *FakeNotifier) Notify(msg string) {
	// if f == nil {
	// 	fmt.Println("f is nil")
	// 	return
	// }
	f.Called = true
}

func main() {
	var f *FakeNotifier = nil
	var n Notifier = f
	fmt.Println(n == nil)
	// below will panic because it won't find the f.Called field
	// For this to not panic implement the if statement in the Notify method
	n.Notify("hi")
}

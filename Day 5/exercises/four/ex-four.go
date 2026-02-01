// Exercise 4 — Single Interface, Multiple Implementations
// Goal: Core interface intuition.
// Define:
// type Notifier interface {
// 	Notify(msg string)
// }
// Create two structs:
// EmailNotifier
// SMSNotifier
// Implement Notify for both.
// Write:
// func Send(n Notifier)
// Call Send with both implementations.

package main

import "fmt"

type Notifier interface {
	Notify(msg string)
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(msg string) {
	fmt.Println(msg)
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Notify(msg string) {
	fmt.Println(msg)
}

func main() {
	emailNotifier := EmailNotifier{Email: "test@test.com"}
	smsNotifier := SMSNotifier{Phone: "1234567890"}
	Send(emailNotifier)
	Send(smsNotifier)
}

func Send(n Notifier) {
	n.Notify("Hello, world!")
}

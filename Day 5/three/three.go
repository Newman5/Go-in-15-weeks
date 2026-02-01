// Define an interface Notifier with method Notify(msg string).
// Create a struct EmailNotifier.
// Implement Notify on it.
// Write a function Send(n Notifier) that calls Notify.

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

func main() {
	emailSending := EmailNotifier{Email: "test@test.com"}
	Send(emailSending)
}

func Send(n Notifier) {
	n.Notify("Email sent!")
}

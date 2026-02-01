// Exercise 5 — Same Struct, Multiple Interfaces
// Goal: Break the “one type → one interface” mental model.
// Define two interfaces:
// type Notifier interface {
// 	Notify(string)
// }
// type Logger interface {
// 	Log(string)
// }
// Create one struct:
// type Service struct{}
// Implement both methods on Service.
// Assign:
// var n Notifier = &Service{}
// var l Logger   = &Service{}
// Call methods through both variables.

package main

import "fmt"

type Notifier interface {
	Notify(string)
}

type Logger interface {
	Log(string)
}

type Delivery struct {
	Address string
}

func (d Delivery) Notify(msg string) {
	fmt.Println(msg)
}

func (d Delivery) Log(msg string) {
	fmt.Println(msg)
}

func main() {
	delivery := Delivery{Address: "123 Main St"}
	delivery.Notify("Package delivered to 123 Main St")
	delivery.Log("Package delivered to 123 Main St")
	// These are used when you don't want the fields in struct to be exposed
	// Eg. in the below n.Address will cause an error
	var n Notifier = &delivery
	var l Logger = &delivery
	n.Notify("Package delivered to 123 Main St via interface")
	l.Log("Package delivered to 123 Main St via interface")
}

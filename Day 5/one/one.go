// Define a struct User with fields Name string and Age int.
// Add a method IsAdult() that returns bool.
// Call it from main.

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) SetAge(age int) {
	u.Age = age
}

func main() {
	bob := User{
		Name: "Bob",
		Age:  20,
	}
	fmt.Println(bob.IsAdult())
	lob := User{
		Name: "Lob",
		Age:  15,
	}
	fmt.Println(lob.IsAdult())
	lob.SetAge(20)
	fmt.Println(lob.Age)
}

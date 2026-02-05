package main

import (
	"fmt"
)

func main() {
	user := User{
		Name: "John",
		Age:  20,
	}
	user1 := user.SetCopyname("Jane")
	fmt.Println(user1)
	fmt.Println(user)

	user.SetName("Jane")
	fmt.Println(user)
}

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := make(map[string]User)
	users["Alice"] = User{Name: "Alice", Age: 20}
	users["Bob"] = User{Name: "Bob", Age: 30}
	users["Charlie"] = User{Name: "Charlie", Age: 40}
	fmt.Println(users["Charlie"])

}

package main

import "fmt"

func main() {
	var counts map[string]int
	fmt.Println(counts)
	counts["a"] = 10
	fmt.Println(counts["a"])

}

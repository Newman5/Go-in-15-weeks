package main

import "fmt"

func main() {
	counts := make(map[string]int)
	counts["a"] = 10
	counts["b"] = 20
	counts["c"] = 30
	counts["d"] = 40
	counts["e"] = 50

	// The order changes each time the for loop runs
	for k, v := range counts {
		fmt.Println(k, v)
	}
	for k, v := range counts {
		fmt.Println(k, v)
	}
	for k, v := range counts {
		fmt.Println(k, v)
	}
	for k, v := range counts {
		fmt.Println(k, v)
	}
}

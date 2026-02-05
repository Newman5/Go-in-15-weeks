package main

import "fmt"

func main() {
	m := make(map[string]int) // This initializes w/ make
	// var m map[string]int
	// m = map[string]int{} // This initialized manually
	fmt.Println(m)
	MapsRead(m)
	m = MapsWrite(m)
	fmt.Println(m)
	MapReadWithKey(m, "a")
	MapReadWithKey(m, "x")
}

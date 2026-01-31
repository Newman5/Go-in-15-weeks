package main

import "fmt"

func main() {
	counts := make(map[string]int)
	counts["a"] = 10
	counts["b"] = 20
	counts["c"] = 30

	i := 10
	changeIntNot(i, 20)
	fmt.Println(i)
	changeIntAndMap(&i, 20, "a", counts)
	fmt.Println(i)
	fmt.Println(counts)

}

func changeIntAndMap(i *int, j int, a string, m map[string]int) {
	// Pointer needed to change ints, strings, bools etc.
	*i = j
	// No pointer needed to change maps
	m[a] = j
}

func changeIntNot(i int, j int) {
	i = j
}

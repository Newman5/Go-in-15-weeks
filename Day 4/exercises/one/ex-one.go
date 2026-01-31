package main

import "fmt"

func main() {
	scores := make(map[string]int)
	scores["Alice"] = 10
	scores["Bob"] = 0

	// fmt.Println(scores["Alice"])
	// fmt.Println(scores["Bob"])
	// fmt.Println(scores["Charlie"])
	// a, ok := scores["Charlie"]
	// fmt.Println(a, ok)

	fmt.Println(getScore(scores, "Alice"))
	fmt.Println(getScore(scores, "Bob"))
	fmt.Println(getScore(scores, "Charlie"))

	addScore(scores, "Alice", 10)
	fmt.Println(getScore(scores, "Alice"))
}

func getScore(scores map[string]int, name string) (int, bool) {
	score, ok := scores[name]
	return score, ok
}

func addScore(scores map[string]int, name string, delta int) {
	scores[name] += delta
}

package main

import "fmt"

func main() {
	names := []string{"a", "b", "a", "c", "b", "a"}
	n := make(map[string]int)
	countChars(names, n)
	for k, v := range n {
		fmt.Println(k, v)
	}
}

func countChars(names []string, m map[string]int) (map[string]int, error) {
	if len(names) == 0 {
		return map[string]int{}, fmt.Errorf("slice is empty")
	}
	for i := range names {
		m[names[i]]++
	}
	return m, nil
}

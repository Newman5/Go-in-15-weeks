package main

import "fmt"

func MapsRead(m map[string]int) {
	for _, k := range m {
		fmt.Println(k)
	}
}

func MapReadWithKey(m map[string]int, key string) {
	value, ok := m[key]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(value)
	}
}

func MapsWrite(m map[string]int) map[string]int {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	return m
}

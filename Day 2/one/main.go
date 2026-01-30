package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 1.0
	y := 23.0
	z := 0.0
	k := 0.1

	fmt.Println(add(x, y))

	result, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result1, err1 := divide(k, z)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}

func add(a float64, b float64) float64 {
	return a + b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}

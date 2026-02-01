// Exercise 1 — Methods vs Functions (warm-up)
// Goal: Internalize why methods exist.
// Define a struct:
//
//	type Temperature struct {
//		Celsius float64
//	}
//
// Write two versions of the same logic:
// A function: ToFahrenheit(t Temperature) float64
// A method: func (t Temperature) ToFahrenheit() float64
// Call both from main.

package main

import "fmt"

type Temperature struct {
	Celsius float64
}

func ToFahrenheit(t Temperature) (float64, error) {
	if t.Celsius < -273.15 {
		return 0, fmt.Errorf("temperature below absolute zero")
	}
	return (t.Celsius*1.8 + 32), nil
}

func (t Temperature) ToFahrenheit() (float64, error) {
	if t.Celsius < -273.15 {
		return 0, fmt.Errorf("temperature below absolute zero")
	}
	return (t.Celsius*1.8 + 32), nil
}

func main() {
	t := Temperature{Celsius: 100}
	a, error1 := ToFahrenheit(t)
	b, error2 := t.ToFahrenheit()

	if error1 != nil || error2 != nil {
		fmt.Println("error")
	} else {
		fmt.Println(a, b)
	}
}

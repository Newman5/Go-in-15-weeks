package utils

import "fmt"

func SumAvg(s []int) (int, float64, error) {
	if len(s) == 0 {
		return 0, 0, fmt.Errorf("slice is empty")
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	avg := float64(sum) / float64(len(s))
	return sum, avg, nil
}

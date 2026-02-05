package utils

import "fmt"

func ModifySlice(s []int, m int, i int) ([]int, error) {
	if s == nil {
		return nil, fmt.Errorf("slice is nil or empty")
	}
	if i < 0 || i >= len(s) {
		return nil, fmt.Errorf("index out of range")
	}
	(s)[i] = m
	return s, nil
}

func AppendSlice(s []int, m int) ([]int, error) {
	if s == nil {
		return nil, fmt.Errorf("slice is nil")
	}
	s = append(s, m)
	return s, nil
}

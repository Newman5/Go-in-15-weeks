package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	count, sum, min, max, err := stats(nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count, sum, min, max)
	}
}

func stats(nums []int) (count int, sum int, min int, max int, err error) {
	if len(nums) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("slice is empty")
	}
	count = len(nums)
	min = nums[0]
	max = nums[0]
	for _, n := range nums {
		sum += n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return count, sum, min, max, nil
}

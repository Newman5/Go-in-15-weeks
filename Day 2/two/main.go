package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums1 := []int{}
	// fmt.Println(len(nums))
	// total, err := addSlice(nums)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(total)
	// }

	min, err := minSlice(nums1)
	max, err1 := maxSlice(nums1)

	if err != nil || err1 != nil {
		fmt.Println("slice is empty, cannot compute min and max")
	} else {
		fmt.Println(min, max)
	}
}

// func addSlice(nums []int) (int, error) {
// 	if len(nums) == 0 {
// 		return 0, fmt.Errorf("slice is empty")
// 	}
// 	total := 0
// 	for _, n := range nums {
// 		total = total + n
// 	}
// 	return total, nil
// }

func minSlice(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min, nil
}

func maxSlice(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max, nil
}

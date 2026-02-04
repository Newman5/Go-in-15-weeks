package main

import (
	"fmt"
	"seven/utils"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums, err := utils.ModifySlice(nums, 10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nums)
	}
	// No change to original slice
	nums1, err1 := utils.AppendSlice(nums, 6)
	if err != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(nums1)
		fmt.Println(nums)
	}

}

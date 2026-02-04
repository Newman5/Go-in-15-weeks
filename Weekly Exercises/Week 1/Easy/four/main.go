package main

import (
	"Week1/Easy/four/utils"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum, avg, err := utils.SumAvg(nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum, avg)
	}
}

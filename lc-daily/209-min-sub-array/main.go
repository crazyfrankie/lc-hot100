package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(nums []int, target int) int {
	var sum, left, right int
	res := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		sum += nums[right]
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if res == math.MaxInt32 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(minSubArrayLen([]int{2, 3, 1, 2, 4, 3}, 7))
}

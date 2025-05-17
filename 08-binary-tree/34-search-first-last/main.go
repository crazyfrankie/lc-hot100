package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // print [3,4]
}

func searchRange(nums []int, target int) []int {
	sch := func(n int, f func(int) bool) int {
		i, j := 0, n
		for i < j {
			h := int(uint(i+j) >> 1)
			if !f(h) {
				i = h + 1
			} else {
				j = h
			}
		}

		return i
	}

	left := sch(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sch(len(nums), func(i int) bool {
		return nums[i] >= target+1
	}) - 1

	return []int{left, right}
}

//func searchRange(nums []int, target int) []int {
//	left := sort.SearchInts(nums, target)
//	if left == len(nums) || nums[left] != target {
//		return []int{-1, -1}
//	}
//	right := sort.SearchInts(nums, target + 1) - 1
//	return []int{left, right}
//}

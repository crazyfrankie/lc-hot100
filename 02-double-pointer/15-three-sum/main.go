/*
题目要求是给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

思路：双指针，固定一个元素后，以它的下一个元素作为左指针，数组末尾作为右指针，转换为两数之和。
关键点在于去重，由于数组是排序后的，有两个地方要去重，
1. 固定的第一个元素，如果当前元素等于它的前一个元素，则跳过当前元素
2. 双指针所构建的序列中：
如果 1 已经作为了结果组里面的一个元素，
对于左指针来说，如果它的下一个元素，也就是 left+1 还是 1，那么直接跳过；
对于右指针来说，如果它的下一个元素，也就是 right-1 还是1，那么直接跳过
*/
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums)) // 输出 [[-1,-1,2],[-1,0,1]]
}

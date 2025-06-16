package main

import "fmt"

// 本题大致起码有两种解法：
// 1. 暴力遍历，两层 for
// 2. 前缀最小值
// 第二种方法具体来说：
// 在第一种方法的基础上，我们需要枚举 j 的值，那么可以借此维护一个前缀最小值，代表从 nums[0……j] 的一个最小值，那么在枚举过程中，
// 如果这个最小值比当前的值要大，那么需要计算差值并和上次的结果进行比较找出最大值，
// 如果这个最小值比当前的值还要小，那么需要进行一个替换，最终得到的 res 就是结果
func maximumDifference(nums []int) int {
	prefMin := nums[0]
	var res int
	for j := 1; j < len(nums); j++ {
		if nums[j] > prefMin {
			res = max(res, nums[j]-prefMin)
		} else {
			prefMin = nums[j]
		}
	}

	return res
}

func main() {
	fmt.Println(maximumDifference([]int{7, 1, 5, 4}))
}

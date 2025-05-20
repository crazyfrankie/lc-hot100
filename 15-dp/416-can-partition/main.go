package main

import "fmt"

// dp[j] 代表从数组中选取若干数是否能凑出 j
func canPartition(nums []int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true // 什么都不取，能凑出 0

	for i := 0; i < len(nums); i++ {
		// 为什么是到 nums[i] ? 而不是 0 ？
		// 因为 j 代表容量，如果 j 小于 nums[i], 那么 nums[i] 压根放不进去，也即这个组合的值不可达，所以不用讨论
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

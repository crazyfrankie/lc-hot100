package main

import "fmt"

// 最长递增子序列：子序列问题
// dp[i] 定义：以 i 为结尾的最长子序列的长度
// 同时本题不像以往的 dp 题目，结果为 dp[len(dp)-1]
// 结合 dp 定义：由于不一定最长子序列的末尾就是 nums 的最后一个元素，可能出现在中间，那么需要在遍历过程中不断的更新一个最大值
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	res := dp[0]
	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

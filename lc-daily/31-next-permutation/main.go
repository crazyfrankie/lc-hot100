package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// 1. 从后向前找第一个下降的位置 i
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 { // 如果不是完全降序（如 [3,2,1]）
		j := n - 1
		// 2. 从后向前找第一个比 nums[i] 大的 nums[j]
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// 3. 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 4. 反转 i+1 到末尾的部分
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}

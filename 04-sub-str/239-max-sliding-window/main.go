package main

import (
	"fmt"
)

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	res := make([]int, 0, len(nums)-k+1)
	q := make([]int, 0) // 存储的是索引

	for i := 0; i < len(nums); i++ {
		// 移除超出窗口范围的元素
		for len(q) > 0 && q[0] <= i-k {
			q = q[1:]
		}

		// 保持队列单调递减
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}

		// 加入当前索引
		q = append(q, i)

		// 记录当前窗口的最大值（窗口大小达到 k 时才开始记录）
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}

	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

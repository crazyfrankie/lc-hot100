package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var max int
	for left < right {
		s := (right - left) * min(height[left], height[right])
		if s > max {
			max = s
		}
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return max
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums)) // 输出49
}

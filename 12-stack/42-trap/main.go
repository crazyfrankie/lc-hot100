package main

import "fmt"

func trap(height []int) int {
	stack := make([]int, 1, len(height))
	res := 0
	for i := 1; i < len(height); i++ {
		if height[stack[len(stack)-1]] >= height[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					left := stack[len(stack)-1]
					area := (min(height[left], height[i]) - height[mid]) * (i - left - 1)
					res += area
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

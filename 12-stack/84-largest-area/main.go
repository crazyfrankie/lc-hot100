package main

import "fmt"

func largestRectangleArea(heights []int) int {
	res := -1
	newHeights := make([]int, 0, len(heights)+2)
	newHeights = append(newHeights, 0)
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)

	stack := make([]int, 0, len(newHeights))
	stack = append(stack, 0)
	for i := 1; i < len(newHeights); i++ {
		if len(stack) > 0 && newHeights[i] >= newHeights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					left := stack[len(stack)-1]
					wide := i - left - 1
					area := wide * newHeights[mid]
					res = max(res, area)
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// 从左到右遍历上层
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上到下遍历右层
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 检查是否还有下层需要遍历
		if top <= bottom {
			// 从右到左遍历下层
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// 检查是否还有左层需要遍历
		if left <= right {
			// 从下到上遍历左层
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

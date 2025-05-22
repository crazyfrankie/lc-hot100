package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int
	var backtrack func(target, start int)
	backtrack = func(target, start int) {
		if len(path) == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(target, i+1)
			path = path[:len(path)-1]
		}
	}

	for i := 0; i <= len(nums); i++ {
		backtrack(i, 0)
	}

	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	var backtrack func()
	used := make([]bool, len(nums))
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()

	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

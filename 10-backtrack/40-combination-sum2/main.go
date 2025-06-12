package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int
	var sum int
	var backtrack func(start int)
	backtrack = func(start int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		} else if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 去重关键点：在同层中跳过重复元素
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i + 1) // 不能重复使用元素，所以 i+1
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)

	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

package main

import "fmt"

// 本题的思路可以参考 77 题 (https://leetcode.cn/problems/combinations/description/)
// 不同的点在于，本题可以取重复的元素，所以 start 不再是 i+1, 而是 i 本身
// 另外本题的终止条件不再是满足长度为 k 的 path，本题不限 path 的长度，而是内部元素中和，所以需要单独维护一个 sum
// 所以在回溯的处理上多了一步 sum -= candidates[i]
// 77 题代码如下，可以参考思路：
func combine(n int, k int) [][]int {
	var result [][]int
	var path []int
	var backtrack func(n, start int)
	backtrack = func(n, start int) {
		if len(path)+(n-start+1) < k {
			return
		}
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(n, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(n, 1)

	return result
}

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	var result [][]int
	var path []int
	var sum int
	var backtrack func(n, start, sum int)
	backtrack = func(n, start, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		} else if sum > target {
			return
		}
		for i := start; i < n; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(n, i, sum)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(n, 0, sum)

	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

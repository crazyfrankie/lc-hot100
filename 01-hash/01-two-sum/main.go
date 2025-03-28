/*
本题要求是给定一个数组和一个目标值，
判断数组中是否存在两个值的和等于目标值，
如果存在则返回它们的下标。

思路：使用哈希，遍历数组，如果 target - n（n代表当前遍历的元素值）存在于 map 中，
则代表这是一对有效结果；如果不存在，将 n 和对应的下标存入 map
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}

	for i, n := range nums {
		if p, ok := numMap[target-n]; ok {
			return []int{p, i}
		}
		numMap[n] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9)) // 输出 [0,1]
}

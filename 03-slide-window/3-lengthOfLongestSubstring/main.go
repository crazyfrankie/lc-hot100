/*
无重复字符的最长字串

滑动窗口思路，采用 map 维护一个窗口
*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left, right := 0, -1
	m := map[byte]int{}
	n := len(s)
	res := 0
	for left < n {
		if left != 0 {
			delete(m, s[left-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		res = max(res, right-left+1)
		left++
	}

	return res
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // 输出 3
}

package main

import "fmt"

func buildPMT(pattern string) []int {
	pmt := make([]int, len(pattern))
	pmt[0] = 0  // 首字符的 PMT 值固定为 0
	length := 0 // 当前最长公共前后缀长度

	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[length] {
			length++
			pmt[i] = length
			i++
		} else {
			if length != 0 {
				length = pmt[length-1] // 回退到前一个匹配位置
			} else {
				pmt[i] = 0
				i++
			}
		}
	}
	return pmt
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	pmt := buildPMT(needle)
	i, j := 0, 0 // i: text指针, j: pattern指针

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j // 匹配成功
			}
		} else {
			if j != 0 {
				j = pmt[j-1] // 利用PMT跳过已匹配部分
			} else {
				i++
			}
		}
	}
	return -1 // 未找到
}

func violentStrStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}

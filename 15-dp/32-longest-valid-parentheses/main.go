package main

import "fmt"

// 这道题我最开始的思路来源于 `有效的括号`, 直接就用栈了, 代码如下
//
//	func longestValidParentheses(s string) int {
//		if len(s) == 0 {
//			return 0
//		}
//
//		stack := make([]byte, 0)
//		res := 0
//		for i := 0; i < len(s); i++ {
//			if s[i] == '(' {
//				stack = append(stack, s[i])
//			} else if len(stack) > 0 {
//				res += 2
//				stack = stack[:len(stack)-1]
//			}
//		}
//
//		return res
//	}
//
// 问题在于没有考虑到题目中的连续, 对于这个用例 "()(()", 输出是 4, 但正确的是 2, 因为这两对是不连续的, 最长只能是 2
// 于是有了下面的解法
// -1 作为哨兵的 双重作用:
//
//  1. 让首个匹配对直接算出正确长度
//     例如 "()"：当 i=1 时弹栈，栈顶是 -1，1 - (-1) = 2；无需额外 +1。
//
//  2. 标记「最后一次失败的位置」
//     如果字符串一开始就是 ')...'，我们会把该 ) 的索引压进去，使后续长度基线自动右移，避免把无效段计入。

// 但官方解法推荐 dp, 个人认为栈更适合, 思路更清晰
func longestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{-1} // 初始值

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 弹出一个
			if len(stack) == 0 {
				stack = append(stack, i) // 重新设定起始点
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestValidParentheses(")(()()"))
}

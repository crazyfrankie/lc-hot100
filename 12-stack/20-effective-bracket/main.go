/*
主要的思路就是要使用 map 记录一下括号有哪些，采用栈的思路，当遍历到左括号（右括号也可，只不过需要改一下 map 的结构），就入栈；
当遍历到右括号时，此时需要查看是否正确闭合，那么就弹栈，栈顶元素和该右括号对应的左括号进行对比
最后如果栈为空代表闭合正确，不为空可能是 s 中压根没有左括号或右括号，如示例的第五个
*/
package main

import "fmt"

func main() {
	strs := []string{"()", "(]", "(]", "([])", "(("}
	for _, s := range strs {
		fmt.Println(isValid(s))
	}
	// print
	// true
	// false
	// false
	// true
	// false
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	str := make([]byte, 0, 3)
	for i := 0; i < len(s); i++ {
		if res, ok := m[s[i]]; ok {
			if len(str) == 0 || res != str[len(str)-1] {
				return false
			}
			str = str[:len(str)-1]
		} else {
			str = append(str, s[i])
		}
	}

	return len(str) == 0
}

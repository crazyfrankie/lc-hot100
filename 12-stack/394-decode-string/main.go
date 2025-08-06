package main

import (
	"fmt"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。
//
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

func decodeString(s string) string {
	var numStack []int
	var strStack []string
	var currentNum int
	var currentStr strings.Builder

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			currentNum = currentNum*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr.String())
			currentNum = 0
			currentStr.Reset()
		} else if ch == ']' {
			// 弹出数字和字符串
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 构建重复后的字符串
			repeatedStr := strings.Repeat(currentStr.String(), repeat)
			currentStr.Reset()
			currentStr.WriteString(prevStr)
			currentStr.WriteString(repeatedStr)
		} else {
			currentStr.WriteRune(ch)
		}
	}

	return currentStr.String()
}

func main() {
	fmt.Println(decodeString("3[a2[c]]") == "accaccacc")
}

package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	var path []byte
	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(path) == n*2 {
			tmp := make([]byte, len(path))
			copy(tmp, path)
			result = append(result, string(tmp))
			return
		}
		if open < n {
			path = append(path, '(')
			backtrack(open+1, close)
			path = path[:len(path)-1]
		}
		if close < open {
			path = append(path, ')')
			backtrack(open, close+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)

	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}

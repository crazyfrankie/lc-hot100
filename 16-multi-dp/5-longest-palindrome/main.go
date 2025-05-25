package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	start, maxLen := 0, -1
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					start = i
				}
			}
		}
	}

	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("abcaba"))
}

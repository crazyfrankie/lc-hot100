package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	find := func(s string) bool {
		for _, word := range wordDict {
			if word == s {
				return true
			}
		}
		return false
	}

	// 先物品再背包：错误！
	// 因为不能假设某个单词就一定存在，在计算 dp[j] 的时候可能 dp[j-len(wordDict[i])] 还没算出来，因为它是根据单词进行移动的
	//for i := 0; i < len(wordDict); i++ {
	//	for j := len(wordDict[i]); j <= n; j++ {
	//		word := s[j-len(wordDict[i]) : len(wordDict[i])]
	//		if find(word) && dp[j-len(wordDict[i])] {
	//			dp[j] = true
	//		}
	//	}
	//}

	// 一定是先背包再物品，要把每个字串作为候选值，挨个的对比，不会出现漏解的情况，不能提前预设“我要用哪个单词贴在哪儿”
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			word := s[j:i]
			if find(word) && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

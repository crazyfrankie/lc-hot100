package main

import "fmt"

// dp[i][j] 的含义是以 i-1 为结尾的 word1 和以 j-1 为结尾的 word2 的最少操作次数

// 递归：
// 如果 word[i-1] 和 word[j-1] 相等：那么当前 dp[i][j] 的值与以 i-2 为结尾的 word1 和以 j-1 为结尾的 word2 的最少操作次数相等，也即 dp[i-1][j-1]
// 如果 word[i-1] 和 word[j-1] 不相等：
//   可以分为三种情况:
//   1. 增, dp[i][j] = dp[i-1][j]+1 或 dp[i][j-1]+1，具体的来说如果当前的两个字符不同，那么即可以对 word1 的上一个字符为结尾的结果之后加上一个字符，也就是 dp[i-1][j]+1，同理对于 word2 来说也可以，那么就是对这两个值取一个最小值
//   2. 删, 删的话其实和增是一样的，比如上面提到可以对 word1 的 i-2 为结尾的进行加字符，我同理可以对 j-1 为结尾的进行减字符对吧，两者是一致的，所以可以算在上面的情况中
//   3. 改, dp[i][j] = dp[i-1][j-1]+1，具体来说，改的话代表只需要在它们两个的前一个字符相同的情况下加一个步骤就行，直接修改成相同就行，那么上一次的结果是 dp[i-1][j-1]，最后加 1 就行。
//   那么最终的结果就是三者取最小值

// 初始化，由于本题的推导都是从左上方和正上方进行推导，那么只需要对第一行和第一列进行初始化
//   - 第一列： dp[i][0], 代表以 -1 为结尾的 word2（即空字符串）， i-1 为结尾的 word1 的最少操作次数，那么最简单的就是将 word1 删除至空字符串，次数与长度相等，即 i
//            那么有 dp[i][0] = i
//   - 第一行： 其实同上， dp[0][j] = j

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

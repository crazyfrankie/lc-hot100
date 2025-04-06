/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/
package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		num := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			num = min(num, dp[i-j*j])
		}
		dp[i] = num + 1
	}
	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	res := numSquares(12)
	fmt.Println(res)
}

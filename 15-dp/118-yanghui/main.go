package main

import "fmt"

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	if numRows <= 1 {
		dp[0][0] = 1
	} else {
		dp[0][0], dp[1][0], dp[1][1] = 1, 1, 1
	}
	for i := 2; i < numRows; i++ {
		length := len(dp[i])
		fmt.Println(length)
		for j := 0; j < length; j++ {
			if j != 0 && j != length-1 {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				continue
			}
			dp[i][j] = 1
		}
	}

	return dp
}

func main() {
	dp := generate(5)
	fmt.Println(dp)
}

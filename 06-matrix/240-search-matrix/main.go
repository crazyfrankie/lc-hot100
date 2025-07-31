package main

import "fmt"

// 直观的思路是 从左到右，从上到下的遍历，但是这样会遇到一些问题
// 在我们遍历到某个元素时，如果它比 target 大，我们有两种走法：往下/往右；第二，如果比它小怎么办？怎么回退？
// 当然这些问题可以解决，只不过比较麻烦，不够高效。
//
// 于是利用这个矩阵的特性，从左到右递增，从上到下递增，那么很容易可以发现，
// 右上角的这个元素，它满足本行最大，本列最小，如果从它开始遍历的话：
// - 如果这个点比 target 大，说明这一列的所有数也比 target 大（因为列递增），所以可以排除这一列 → 左移
// - 如果这个点比 target 小，说明这一行的所有数也比 target 小（因为行递增），所以可以排除这一行 → 下移
// 这样的话遍历的逻辑就很清晰了，同样的，从右下角的元素开始遍历也可以，思路相同。
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col-- // move left
		} else {
			row++ // move down
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
}

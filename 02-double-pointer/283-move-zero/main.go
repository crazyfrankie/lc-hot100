/*
题目要求是给定一个整数数组，将数组中的所有 0 移动到数组末尾，并不改变其他元素的相对位置，且是原地修改。

思路：采用双指针，
左指针的指向代表当前要被覆盖的元素，右指针进行遍历，如果未指向 0，则将其赋值给左指针所指向的位置，然后将右指针的位置赋值为 0。
[0,1,0,3,12]
第一次循环，left = 0, right = 1, 此时 nums[right] != 0, 则进行 nums[left] = nums[right], nums[right] = 0
结果为
[1,0,0,3,12]
后面的以此类推
*/
package main

import "fmt"

func moveZeroes(nums []int) {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left] = nums[right]
			if right != 0 {
				nums[right] = 0
			}
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // 输出 [1,3,12,0,0,0]
}

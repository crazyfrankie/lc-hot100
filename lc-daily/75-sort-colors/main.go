package main

import "fmt"

// 这个问题是经典的荷兰国旗问题，我们需要将数组中的0、1、2按照顺序排列
// 指针定义：
//  p0：指向下一个0应该放置的位置
//  p2：指向下一个2应该放置的位置
//  curr：当前遍历的指针
// 遍历过程：
//  当遇到0时，与p0位置的元素交换，并移动p0和curr
//  当遇到1时，直接跳过（已经在正确位置）
//  当遇到2时，与p2位置的元素交换，并移动p2（不移动curr，因为交换过来的元素需要再次检查）
// 终止条件：当curr超过p2时，所有元素已经处理完毕
func sortColors(nums []int) {
	// 初始化三个指针：
	// p0: 指向0的最右边界
	// p2: 指向2的最左边界
	// curr: 当前遍历的元素
	p0, curr, p2 := 0, 0, len(nums)-1

	for curr <= p2 {
		switch nums[curr] {
		case 0:
			// 交换当前元素和p0位置的元素
			nums[curr], nums[p0] = nums[p0], nums[curr]
			p0++
			curr++
		case 1:
			// 1已经在正确位置，只需移动curr
			curr++
		case 2:
			// 交换当前元素和p2位置的元素
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
			// 这里不移动curr，因为交换过来的元素需要再次检查
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

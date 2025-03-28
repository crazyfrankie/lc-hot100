/*
本题关键点在于要先固定一个元素 x，
然后进行递增，判断 x+1 x+2 x+3 .etc 是否存在于数组中，在这个过程中不断更新最长长度。
那么其实两层 for 就能解决，但细想的话，这种思路我们做了很多无用功。
比如 当你固定了元素 x，并判断了 x+1 x+2  是否存在于数组中，此时我们假设 x+1 是存在的，
当你下次选中 x+1 作为 x 时，你又会判断一次 x+2，那么我们可以使用 map 做一次筛选，减少判断次数。
具体的来说，先把所有元素存到 map 中，当要选中了一个元素作为 x 时，先判断 x-1 是否存在于 map 中，
如果存在，代表这个元素及之后的序列已经被判断过了，不需要再判断;
如果不存在，才进行 x+1 x+2 的判断，以及更新最长长度。
*/

package main

import "fmt"

func longestConsecutive(nums []int) int {
	max := 0
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}

	for num := range numMap {
		if !numMap[num-1] {
			curr := num
			currL := 1
			for numMap[curr+1] {
				curr++
				currL++
			}
			if max < currL {
				max = currL
			}
		}
	}

	return max
}

func main() {
	nums := []int{100, 4, 200, 1, 2, 3}
	fmt.Println(longestConsecutive(nums)) // 输出4
}

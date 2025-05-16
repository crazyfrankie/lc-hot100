/*
主要是一个堆性质的应用，但重点在于，对数组进行堆化后，数组不一定是严格递增或递减，只是局部满足堆性质，
那么需要根据 k 的值，将堆顶的元素逐个移到数组末尾，然后再堆化，以达到每次移动的元素都是当前最大的
*/
package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapify := func(nums []int, i, n int) {
		for {
			left := 2*i + 1
			right := 2*i + 2
			candidate := i

			if left < n && nums[candidate] < nums[left] {
				candidate = left
			}
			if right < n && nums[candidate] < nums[right] {
				candidate = right
			}
			if candidate == i {
				break
			}
			nums[i], nums[candidate] = nums[candidate], nums[i]
			i = candidate
		}
	}

	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	for i := n - 1; i >= n-k; i-- {
		// 将当前最大元素(堆顶)移到数组末尾
		nums[0], nums[i] = nums[i], nums[0]
		// 调整剩余堆
		heapify(nums, 0, i)
	}

	return nums[n-k]
}

func main() {
	nums := [][]int{{3, 2, 3, 1, 2, 4, 5, 5, 6}, {3, 2, 1, 5, 6, 4}}
	fmt.Println(findKthLargest(nums[0], 4)) // print 4
	fmt.Println(findKthLargest(nums[1], 2)) // print 5
}

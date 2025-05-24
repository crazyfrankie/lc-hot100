package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 堆排
//
//	func sortArray(nums []int) []int {
//		n := len(nums)
//		heapify := func(nums []int, i, n int) {
//			for {
//				left := 2*i + 1
//				right := 2*i + 2
//				candidate := i
//
//				if left < n && nums[left] > nums[candidate] {
//					candidate = left
//				}
//				if right < n && nums[right] > nums[candidate] {
//					candidate = right
//				}
//				if i == candidate {
//					break
//				}
//				nums[i], nums[candidate] = nums[candidate], nums[i]
//				i = candidate
//			}
//		}
//
//		for i := n/2 - 1; i >= 0; i-- {
//			heapify(nums, i, n)
//		}
//
//		for i := n - 1; i >= 0; i-- {
//			nums[0], nums[i] = nums[i], nums[0]
//			heapify(nums, 0, i)
//		}
//
//		return nums
//	}

func sortArray(nums []int) []int {
	rand.Seed(time.Now().UnixNano())

	var partition func([]int, int, int) int
	partition = func(nums []int, l, r int) int {
		pivot := nums[r]
		i := l - 1
		for j := l; j <= r-1; j++ {
			if nums[j] < pivot {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[r] = nums[r], nums[i+1]
		return i + 1
	}

	var randomizedPartition func([]int, int, int) int
	randomizedPartition = func(nums []int, l, r int) int {
		i := rand.Intn(r-l+1) + l
		nums[i], nums[r] = nums[r], nums[i]
		return partition(nums, l, r)
	}

	var randomizedQuickSort func([]int, int, int)
	randomizedQuickSort = func(nums []int, l, r int) {
		if l < r {
			pos := randomizedPartition(nums, l, r)
			randomizedQuickSort(nums, l, pos-1)
			randomizedQuickSort(nums, pos+1, r)
		}
	}

	randomizedQuickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{4, 6, 8, 5, 9}))
}

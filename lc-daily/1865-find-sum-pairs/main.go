package main

import "fmt"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt   map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	cnt := map[int]int{}
	for _, x := range nums2 {
		cnt[x]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, cnt: cnt}
}

func (p *FindSumPairs) Add(index int, val int) {
	p.cnt[p.nums2[index]]--
	p.nums2[index] += val
	p.cnt[p.nums2[index]]++
}

func (p *FindSumPairs) Count(tot int) int {
	var res int
	for _, x := range p.nums1 {
		res += p.cnt[tot-x]
	}
	return res
}

func main() {
	t := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(t.Count(7))
	t.Add(3, 2)
	fmt.Println(t.Count(8))
	fmt.Println(t.Count(4))
	t.Add(0, 1)
	t.Add(1, 1)
	fmt.Println(t.Count(7))
}

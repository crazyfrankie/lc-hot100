/*
O(1) 时间插入、删除和随机获取元素

其实这题很简单，首先O(1) 插入和删除，直接就能想到 map, 然后它还包含一个随机，那么就还要一个 rand，但怎么生成随机数呢？
因为你 map 再怎么样，也无法用随机数生成一个 key 来，你又不知道它的 key 究竟有哪些。

那这个时候考虑引入数组, 它真正的存储元素, 在随机数生成时根据它当前的长度去生成。
而 map 由于要达到所有操作都是 O(1), 那么它的 key 和 val 分别是 元素的值和元素的索引，这样就能达到 O(1) 的插入删除和随机访问了
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	r     *rand.Rand
	cache map[int]int // 记录索引
	list  []int       // 存数据
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
		cache: make(map[int]int),
		list:  make([]int, 0),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.cache[val]; ok {
		return false
	}
	r.list = append(r.list, val)
	r.cache[val] = len(r.list) - 1

	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	idx, ok := r.cache[val]
	if !ok {
		return false
	}

	// 将要删除的元素与最后一个元素交换位置
	lastIndex := len(r.list) - 1
	lastVal := r.list[lastIndex]
	r.list[idx] = lastVal  // 覆盖要删除的值
	r.cache[lastVal] = idx // 更新最后一个值在 map 中的索引

	r.list = r.list[:lastIndex]
	delete(r.cache, val)

	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.list[r.r.Intn(len(r.list))]
}

func main() {
	commands := []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}
	values := [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}}

	// 存储输出结果
	var res []interface{}

	// 初始化 RandomizedSet
	var rs *RandomizedSet

	for i, op := range commands {
		switch op {
		case "RandomizedSet":
			rs = NewRandomizedSet()
			res = append(res, nil) // "RandomizedSet" 这个操作没有返回值
		case "insert":
			res = append(res, rs.Insert(values[i][0]))
		case "remove":
			res = append(res, rs.Remove(values[i][0]))
		case "getRandom":
			res = append(res, rs.GetRandom())
		}
	}

	// 打印结果
	fmt.Println(res) // 预期输出: [nil, true, false, true, 2, true, false, 2]
}

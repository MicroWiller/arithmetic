package head

import "container/heap"

// 扩展![](https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/CaseOfHeap.png)

// 347. 前 K 个高频元素：https://leetcode-cn.com/problems/top-k-frequent-elements/description/
// 练习题 1：给定一个数组，求这个数组的前 k 个高频元素。如果有两个数出现次数一样，那么优先取较小的那个数。
// 输入：A = [1,2,1,1,3,3,2,3] k = 2
// 输出：[1, 3]
// 解释：每个数字的出现频率从高到低排序<1, 3次>, <3,3次> <2, 2次>，取前 2 个高频数字。所以是 [1, 3]。
func TopKFrequent(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 || k <= 0 {
		return nil
	}
	// 把元素组装成节点，放入小根堆
	h := new(TopFreqHeap)
	cntMap := make(map[int]int)
	for _, num := range nums {
		if cnt, ok := cntMap[num]; ok {
			cnt++
			cntMap[num] = cnt
		} else {
			cntMap[num] = 1
		}
		newNode := FreqNode{
			key: num,
			cnt: cntMap[num],
		}
		h.Push(newNode)
	}
	heap.Init(h)
	// 弹出 len(nums) - k 次
	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(h)
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		getNode := heap.Pop(h)
		res = append(res, getNode.(FreqNode).key)
	}
	return res
}

type FreqNode struct {
	key int
	cnt int
}

type TopFreqHeap []FreqNode

// Less 小根堆
func (h *TopFreqHeap) Less(i, j int) bool {
	return (*h)[i].cnt < (*h)[j].cnt
}

func (h *TopFreqHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *TopFreqHeap) Len() int {
	return len(*h)
}

func (h *TopFreqHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *TopFreqHeap) Push(v interface{}) {
	*h = append(*h, v.(FreqNode))
}

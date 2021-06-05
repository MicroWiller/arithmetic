package head

// Element 可以定义成其他类型/接口，为了方便直接用int
type Element int

// BigRootHeap ...
type BigRootHeap []Element

// NewBigRootHeap BigRootHeap的创建工厂
func NewBigRootHeap() *BigRootHeap {
	return &BigRootHeap{}
}

// todo: 实现堆的 push 和 pop

// Push 添加元素
// 新元素添加到尾部，并往上浮
func (h *BigRootHeap) Push(v interface{}) {
	*h = append(*h, v.(Element))
	h.swim(h.Len() - 1)
}

// Pop 弹出头节点
func (h *BigRootHeap) Pop() Element {
	rt := (*h)[0]
	// 尾元素放到头节点，并下沉
	(*h)[0] = (*h)[h.Len()-1]
	// 切片长度
	*h = (*h)[:h.Len()-1]
	h.sink(0)
	return rt
}

// Len ...
func (h *BigRootHeap) Len() int {
	return len(*h)
}

// sink 下标i 下沉
func (h *BigRootHeap) sink(i int) {
	for {
		// left child
		j := (i << 1) + 1
		// 校验下标，子节点下标 需要小于总长度，不然会下标越界
		if j >= len(*h) || j < 0 {
			break
		}
		if j+1 < len(*h) && (*h)[j] < (*h)[j+1] {
			// right child
			j = j + 1
		}
		// 如果子节点最大值小于当前节点，退出
		if (*h)[j] < (*h)[i] {
			break
		}
		(*h)[j], (*h)[i] = (*h)[i], (*h)[j]
		i = j
	}
}

func (h *BigRootHeap) swim(i int) {
	for {
		// parent node
		j := (i - 1) / 2
		if j < 0 || j > len(*h) || j == i {
			break
		}
		// 当前节点小于父节点，退出
		if (*h)[j] > (*h)[i] {
			break
		}
		(*h)[j], (*h)[i] = (*h)[i], (*h)[j]
		i = j
	}
}

// 例 1：最小的 k 个数
//【题目】给定一个数组 a[]，返回这个数组中最小的 k 个数。
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
// 输入：A = [3,2,1], k = 2
// 输出：[2, 1]
func GetLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 || k <= 0 {
		return nil
	}
	brh := NewBigRootHeap()
	for i := 0; i < len(arr); i++ {
		brh.Push(Element(arr[i]))
		if brh.Len() > k {
			brh.Pop()
		}
	}
	res := []int{}
	for i := 0; i < brh.Len(); i++ {
		res = append(res, int((*brh)[i]))
	}
	return res
}

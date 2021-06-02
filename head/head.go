package head

// Element 可以定义成类型，为了方便直接用int
type Element int

// BigRootHeap ...
type BigRootHeap []Element

// NewBigRootHeap BigRootHeap的创建工厂
func NewBigRootHeap() *BigRootHeap {
	return &BigRootHeap{}
}

// todo: 实现堆的 push 和 pop

// Push 添加元素
func (h *BigRootHeap) Push(v interface{}) {

}

// Pop ...
func (h *BigRootHeap) Pop() {

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
		// 校验下标
		if j > len(*h) || j < 0 {
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
		// parant node
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

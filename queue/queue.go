package queue

// 二叉树
type TreeNode struct {
	value       int
	left, right *TreeNode
}

// 题目：从上到下按层打印二叉树，同一层结点按从左到右的顺序打印，每一层打印到一行。
func PrintBinaryTree(head *TreeNode) [][]int {
	if head == nil {
		return nil
	}
	// todo：使用切片实现队列
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	// check：添加是否添加到切片最后
	queue = append(queue, head)
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			// check：是否移除了第一个元素
			queue = append(queue[:0], queue[1:]...)
			level = append(level, curNode.value)
			if curNode.left != nil {
				queue = append(queue, curNode.left)
			}
			if curNode.right != nil {
				queue = append(queue, curNode.right)
			}
		}
		res = append(res, level)
	}
	return res
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
type Node struct {
	value             int
	left, right, next *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			// 移除头节点
			queue = append(queue[:0], queue[1:]...)
			if i < size-1 {
				curNode.next = queue[0]
			}
			if curNode.left != nil {
				queue = append(queue, curNode.left)
			}
			if curNode.right != nil {
				queue = append(queue, curNode.right)
			}
		}
	}
	return root
}

// 循环队列：
// index = i 的后一个下标为：(i + 1) % capacity
// index = i 的前一个下标为：(i - 1 + capacity) % capacity
type ElementType int

// ![](https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/CircularQueueByPlus.png)
type CircularQueueByPlus struct {
	capacity int
	front    int
	rear     int // 指向下个将存储的位置
	sizes    []ElementType
}

func NewCircularQueueByPlus(capacity int) *CircularQueueByPlus {
	return &CircularQueueByPlus{
		capacity: capacity,
		front:    0,
		rear:     0,
		sizes:    make([]ElementType, capacity+1),
	}
}

func (c *CircularQueueByPlus) IsEmpty() bool {
	return c.front == c.rear
}

func (c *CircularQueueByPlus) IsFull() bool {
	return (c.rear+1)%c.capacity == c.front
}

func (c *CircularQueueByPlus) EnQueue(value ElementType) bool {
	if c.IsFull() {
		return false
	}
	c.sizes[c.rear] = value
	c.rear = (c.rear + 1) % c.capacity
	return true
}

func (c *CircularQueueByPlus) DeQueue() bool {
	if c.IsEmpty() {
		return false
	}
	// 出队后，front 往后移
	c.front = (c.front + 1) % c.capacity
	return true
}

func (c *CircularQueueByPlus) Front() ElementType {
	if c.IsEmpty() {
		return ElementType(-1)
	}
	return c.sizes[c.front]
}

func (c *CircularQueueByPlus) Rear() ElementType {
	if c.IsEmpty() {
		return ElementType(-1)
	}
	tail := (c.rear - 1 + c.capacity) % c.capacity
	return c.sizes[tail]
}

// 示意图：![](https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/CircularQueue.png)
type CircularQueueByUsed struct {
	capacity int
	used     int
	front    int
	rear     int // 指向下个将存储的位置
	sizes    []ElementType
}

func NewCircularQueueByUsed(capacity int) *CircularQueueByUsed {
	return &CircularQueueByUsed{
		capacity: capacity,
		used:     0,
		front:    0,
		rear:     0,
		sizes:    make([]ElementType, capacity),
	}
}

func (c *CircularQueueByUsed) IsEmpty() bool {
	return c.used == 0
}

func (c *CircularQueueByUsed) IsFull() bool {
	return c.used == c.capacity
}

// 将value放到队列，成功返回true
func (c *CircularQueueByUsed) EnQueue(value ElementType) bool {
	if c.IsFull() {
		return false
	}
	// 如果没有放满，用 sizes[rear] 存放新进来的元素
	c.sizes[c.rear] = value
	// 更新rear
	c.rear = (c.rear + 1) % c.capacity
	c.used++
	return true
}

// 删除队首元素，成功返回true
func (c *CircularQueueByUsed) Dequeue() bool {
	if c.IsEmpty() {
		return false
	}
	// 队首元素往后移
	c.front = (c.front + 1) % c.capacity
	c.used--
	return true
}

// 获得队首元素
func (c *CircularQueueByUsed) Front() ElementType {
	if c.IsEmpty() {
		return ElementType(-1)
	}
	return c.sizes[c.front]
}

// 获得队尾元素
func (c *CircularQueueByUsed) Rear() ElementType {
	if c.IsEmpty() {
		return ElementType(-1)
	}
	beforeRear := (c.rear - 1 + c.capacity) % c.capacity
	return c.sizes[beforeRear]
}

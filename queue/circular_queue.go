package queue

// 循环队列：
// index = i 的后一个下标为：(i + 1) % capacity
// index = i 的前一个下标为：(i - 1 + capacity) % capacity
type ElementType int

// 在多线程编程里面，控制变量越少，越容易实现无锁编程
// ![](https://cdn.jsdelivr.net/gh/MicroWiller/photobed@master/CircularQueueByPlus.png)
type CircularQueueByPlus struct {
	capacity int
	front    int
	rear     int // 指向下个将存储的位置
	sizes    []ElementType
}

func NewCircularQueueByPlus(capacity int) *CircularQueueByPlus {
	return &CircularQueueByPlus{
		capacity: capacity + 1,
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

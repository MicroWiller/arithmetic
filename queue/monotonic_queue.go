package queue

import "container/list"

// 单调递减队列：
// 1. 入队，扩张单调队列的覆盖范围；
// 2. 出队，控制单调递减队列的覆盖范围；
// 3. 队首元素就是覆盖范围的最大值；
// 4. 队列中的元素个数小于覆盖范围元素的个数。

// MonotonicQueueReduce 非严格单调递减队列
type MonotonicQueueReduce struct {
	queue *list.List // 使用双端队列来实现
}

func NewMonotonicQueueReduce() *MonotonicQueueReduce {
	return &MonotonicQueueReduce{
		queue: list.New(),
	}
}

// Push 后进
func (m *MonotonicQueueReduce) Push(val int) {
	for m.queue.Len() != 0 && m.queue.Back().Value.(int) < val {
		m.queue.Remove(m.queue.Back())
	}
	m.queue.PushBack(val)
}

// Pop 出队时候，相等才出队
func (m *MonotonicQueueReduce) Pop(val int) {
	if m.queue.Len() != 0 && m.queue.Front().Value.(int) == val {
		m.queue.Remove(m.queue.Front())
	}
}

// 例 3：滑动窗口的最大值
// 【题目】给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。
// 输入：nums = [1,3,-1,-3,5,3], k = 3
// 输出：[3,3,5,5]
// [239] 滑动窗口最大值：https://leetcode-cn.com/problems/sliding-window-maximum/description/
func MaxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return nil
	}
	m := NewMonotonicQueueReduce()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m.Push(nums[i])
		if i < k-1 {
			continue
		}
		res = append(res, m.queue.Front().Value.(int))
		m.Pop(nums[i-k+1])
	}
	return res
}

// 扩展： 用循环队列来模拟单调队列，求解滑动窗口最大值的题目

// 扩展：求滑动窗口最小值
//【题目】给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最小值。
// 输入：nums = [1,3,-1,-3,5,3], k = 3
// 输出：[-1,-3,-3,-3]
func MinSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	queue := NewMonotonicQueueIncrease()
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		queue.Push(nums[i])
		if i < k-1 {
			continue
		}
		res = append(res, queue.queue.Front().Value.(int))
		queue.Pop(nums[i-k+1])
	}
	return res
}

// 单调递增队列
// 队首元素, 是覆盖范围的最小值
type MonotonicQueueIncrease struct {
	queue *list.List
}

// 当 val <= 队列的头节点，加入队头
func (m *MonotonicQueueIncrease) Push(val int) {
	if m.queue.Len() != 0 &&
		val > m.queue.Front().Value.(int) {
		return
	}
	m.queue.PushFront(val)
}

// pop 移除队尾
func (m *MonotonicQueueIncrease) Pop(val int) {
	if m.queue.Len() != 0 && m.queue.Back().Value.(int) == val {
		m.queue.Remove(m.queue.Back())
	}
}
func NewMonotonicQueueIncrease() *MonotonicQueueIncrease {
	return &MonotonicQueueIncrease{list.New()}
}
